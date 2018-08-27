package main

import (
	"fmt"
)

func main() {
	// make an unbuffered channel
	c := make(chan int)

	// create func that pushes ints to the channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// When we're done putting values on channel, close it
		// makes channel read only
		// The closer gives the range clause a heads up when the
		// channel is done
		close(c)
	}()

	// the range blocks the main loop during running
	// eliminating the need for a waitgroup or sleep command
	for n := range c {
		fmt.Println(n)
	}
}
