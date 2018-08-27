package main

import (
	"fmt"
)

// example of making a bunch of routines
func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	// make a bunch og goroutines
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	// close a bunch of goroutines
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
