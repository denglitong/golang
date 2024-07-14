// Go's approach to concurrency differs from the traditional user of threads and shared memory.
// Philosophically, it can be summarized:
// 	Don't communicate by sharing memory; share memory by communicating.
// Channels allow you to pass references to data structures between goroutines.
// If you consider this as passing the ownership of the data (ability to read and write it),
// they become a powerful and expressive synchronization mechanism.
// Reference: https://go.dev/doc/codewalk/sharemem/
package main

import (
	"log"
	"net/http"
	"time"
)

const (
	numPollers     = 2
	pollInterval   = 60 * time.Second
	statusInterval = 10 * time.Second
	errTimeout     = 10 * time.Second
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

// State represents the last-known state of a URL.
type State struct {
	url    string
	status string
}

// StateMonitor maintains a map that stores the state of the URLs begin polled,
// and prints the current state every updateInterval nanoseconds.
// It returns a chan State to which resource state should be sent.
func StateMonitor(updateInterval time.Duration) chan<- State {
	// create a channel to communicate,
	// write outside this function, read in this function (as log out and monitor data)
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatus)
			case s := <-updates:
				urlStatus[s.url] = s.status
			}
		}
	}()
	return updates
}

// logState prints a state map
func logState(s map[string]string) {
	log.Println("Current states:")
	for k, v := range s {
		log.Printf("%s %s\n", k, v)
	}
}

// Resource represents an HTTP URL to be polled by this program
type Resource struct {
	url      string
	errCount int
}

// Poll executes an HTTP HEAD request for url
// and returns the HTTP status string or an error string
func (r *Resource) Poll() string {
	resp, err := http.Head(r.url)
	if err != nil {
		log.Println("Error", r.url, err)
		r.errCount++
		return err.Error()
	}
	r.errCount = 0
	return resp.Status
}

// SleepThenPushTo sleeps for an appropriate interval (dependent on error state)
// before sending the Response to `done`.
func (r *Resource) SleepThenPushTo(done chan<- *Resource) {
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	done <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	// loop channel until it's closed, it'll block if channel has no value and is not closed.
	for r := range in {
		s := r.Poll()
		status <- State{r.url, s}
		out <- r
	}
}

// ShowCommunicateBySharingMemory show sharing memory usage by using Go concurrency primitives.
func ShowCommunicateBySharingMemory() {
	// Create our input and output channels
	pending, complete := make(chan *Resource), make(chan *Resource)

	// Launch the StateMonitor
	status := StateMonitor(statusInterval)

	// Launch some Poller goroutines
	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}

	// Send some Resources to the pending queue
	go func() {
		for _, url := range urls {
			pending <- &Resource{url: url}
		}
	}()

	for r := range complete {
		go r.SleepThenPushTo(pending)
	}
}
