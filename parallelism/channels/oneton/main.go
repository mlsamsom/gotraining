package main

import (
	"fmt"
)

// many funcs pulling from one channel
func main() {
	c := make(chan int)
	done := make(chan bool)

	// make a channel
	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	// 2 goroutines that access the channel
	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()
	<-done
	<-done
}
