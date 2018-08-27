package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // create an unbuffered channel

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // this stops until something takes the value off
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	// Sleep to let the code finish
	// Usually use a waitgroup for this
	time.Sleep(time.Second)
}
