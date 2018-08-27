// N functions writing to the same channel

package main

import (
	"fmt"
	"sync"
)

func main() {
	// race()
	norace()
}

// This code includes a race condition
func race() {
	c := make(chan int)

	var wg sync.WaitGroup

	// multiple go routines are trying to access the shared variable wg
	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

// no race
func norace() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// multiple go routines are trying to access the shared variable wg
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done() // sends message that goroutine is done
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		// waits for 2 done signals
		wg.Wait()
		// then closes
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
