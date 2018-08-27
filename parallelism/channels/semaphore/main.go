// semaphore is a variable that is changed to control
// access to some system resource.
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// use a channel semaphore in place of the WaitGroup type
	done := make(chan bool)

	// a couple goroutines adding ints to the channel c
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// this sits and waits until a couple of dones have come through
	// note that this operation needs to be in a separate goroutine
	// if this is in the main loop it'll block the execution so it
	// never hits the range clause
	go func() {
		<-done // this syntax just throws the value away
		<-done
		close(c)
	}()

	// the range blocks until something is send from c
	for n := range c {
		fmt.Println(n)
	}

}
