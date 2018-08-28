package main

import (
	"fmt"
)

// A series of stages connected by channels
// a chain of functions that take and consume channels
//  a lot like a computation graph, start: inbound channel,
// hidden channels and an outbound channed
func main() {
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
}

// input channel
// consumes an arbitrary number of integers and places them on a channel
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// output channel
// squares everything in the input channel
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
