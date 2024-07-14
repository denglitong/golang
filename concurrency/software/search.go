package software

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) Result

type Result struct {
	str string
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{fmt.Sprintf("%s result for %q", kind, query)}
	}
}

func SequenceSearch(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func ConcurrentSearch(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()
	for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}
	return
}

func ConcurrentSearchWithTimeout(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(50 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}

	}
	return
}

// First Send requests to multiple replicas, and use the first response.
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

var (
	Web1   = fakeSearch("web1")
	Web2   = fakeSearch("web2")
	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2")
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
)

func ConcurrentSearchAvoidTimeoutReduceTailLatency(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()

	timeout := time.After(50 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}

	return
}

func ShowSequenceSearch() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	var results []Result
	results = SequenceSearch("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, "SequenceSearch")
}

func ShowConcurrentSearch() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	var results []Result
	results = ConcurrentSearch("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, "ConcurrentSearch")
}

func ShowConcurrentSearchWithTimeout() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	var results []Result
	results = ConcurrentSearchWithTimeout("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, "ConcurrentSearchWithTimeout")
}

func ShowConcurrentSearchAvoidTimeoutReduceTailLatency() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	var results []Result
	results = ConcurrentSearchAvoidTimeoutReduceTailLatency("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, "ConcurrentSearchAvoidTimeoutReduceTailLatency")
}
