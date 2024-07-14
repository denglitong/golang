package player

import (
	"fmt"
	"math/rand"
	"time"
)

func ShowGoroutine() {
	go boring("boring!")
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("ShowGoroutine: You're boring; I'm leaving.")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func ShowChannel() {
	c := make(chan string)
	go boringWithChan("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("ShowChannel: You're boring; I'm leaving.")
}

func boringWithChan(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func ShowChannelPattern() {
	c := boringReturnChan("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("ShowChannelPattern: You're boring; I'm leaving.")
}

func boringReturnChan(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	// return the channel to the caller
	return c
}

func ShowChannelAsHandlerOnService() {
	joe := boringReturnChan("Joe")
	ann := boringReturnChan("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("ShowChannelAsHandlerOnService: You're both boring; I'm leaving.")
}

func ShowChannelFanInAndMultiplexing() {
	c := fanIn(boringReturnChan("Joe"), boringReturnChan("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("ShowChannelFanInAndMultiplexing: You're both boring; I'm leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

type Message struct {
	str  string
	wait chan bool
}

func ShowRestoringSequences() {
	waitForIt := make(chan bool)
	joe := boringReturnMessageChan("Joe", waitForIt)
	ann := boringReturnMessageChan("Ann", waitForIt)
	for i := 0; i < 5; i++ {
		msg1 := <-joe
		fmt.Println(msg1.str)

		msg2 := <-ann
		fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("ShowRestoringSequences: You're both boring; I'm leaving.")
}

func boringReturnMessageChan(msg string, waitForIt chan bool) <-chan Message {
	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

func ShowChannelFanInAndMultiplexingWithSelect() {
	c := fanIn(boringReturnChan("Joe"), boringReturnChan("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("ShowChannelFanInAndMultiplexingWithSelect: You're both boring; I'm leaving.")
}

func fanInWithSelect(input1, input2 chan string) <-chan string {
	c := make(chan string)
	go func() {
		select {
		case s := <-input1:
			c <- s
		case s := <-input2:
			c <- s
		}
	}()
	return c
}

func ShowBranchTimeoutUsingSelect() {
	c := boringReturnChan5Items("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			// timeout for select branch,
			// if all branches blocked then after 1s this branch will be active
			fmt.Println("ShowBranchTimeoutUsingSelect: You're too slow.")
			return
		}
	}
}

func boringReturnChan5Items(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	// return the channel to the caller
	return c
}

func ShowForLoopTimeoutUsingSelect() {
	c := boringReturnChan("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("ShowForLoopTimeoutUsingSelect: You talk too much.")
			return
		}
	}
}

func ShowChannelQuit() {
	quit := make(chan bool)
	c := boringWithChanQuit("Joe", quit)
	rand.Seed(time.Now().UnixNano())
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Println("ShowChannelQuit: done.")
}

func boringWithChanQuit(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%q %d", msg, i):
				//
			case <-quit:
				return
			}
		}
	}()
	return c
}

func ShowChanQuitReceiveCleanUp() {
	quit := make(chan string)
	c := boringWithChanQuitCleanUp("Joe", quit)
	rand.Seed(time.Now().UnixNano())
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Kill 9"
	fmt.Println("ShowChannelQuit:", <-quit)
}

func boringWithChanQuitCleanUp(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%q %d", msg, i):
				//
			case s := <-quit:
				fmt.Println("Quit begin:", s)
				cleanup()
				// ensure the cleanup finish then we allow to quit
				// the caller need to receiver this message
				quit <- "Quit done."
				return
			}
		}
	}()
	return c
}

func cleanup() {
	fmt.Println("Cleaning up...")
	time.Sleep(2 * time.Second)
	fmt.Println("Cleaning done...")
}

func ShowDaisyChain() {
	const n = 10000
	leftmost := make(chan int)
	var left, right chan int
	left = leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go shift(left, right)
		left = right
	}
	go func() {
		right <- 1
	}()
	fmt.Println("ShowDaisyChain:", <-leftmost)
}

func shift(left, right chan int) {
	left <- 1 + <-right
}
