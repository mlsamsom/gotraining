package main

import (
	"fmt"
)

func main() {
	// deadlock()
	// fixed()
	// loopError()
	fixedLoopError()
}

// why is this a deadlock"
// - 1 is being sent to the channel which is on the main goroutine
// - there is no separate thread to recieve the value
func deadlock() {
	c := make(chan int)
	c <- 1 // this blocks waiting for something to consume the channel
	fmt.Println(<-c)
}

func fixed() {
	c := make(chan int)
	// make new routine
	go func() {
		c <- 1 //this blocks in a separate routine allowing the below to execute
	}()
	fmt.Println(<-c)
}

// this only prints zero
func loopError() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	fmt.Println(<-c)
}

func fixedLoopError() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
