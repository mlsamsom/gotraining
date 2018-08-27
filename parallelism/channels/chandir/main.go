package main

import (
	"fmt"
)

func main() {
	// makes channel c
	c := incrementor()
	// clears channel c and makes cSum
	cSum := puller(c)
	// clears cSum and blocks until done
	for n := range cSum {
		fmt.Println(n)
	}
}

// Makes a channel that recieves incrementing values
// the arrow signifies that the channel can be recieved from only
func incrementor() <-chan int {
	out := make(chan int)
	// put values on the channel in separate goroutine
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// c is a recieve only channel
func puller(c <-chan int) <-chan int {
	out := make(chan int)
	// accumulate values that got put on c
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
