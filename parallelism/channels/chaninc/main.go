package main

import (
	"fmt"
)

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter:", <-c3+<-c4) // these receipts block
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		// loop 20 times and push 1 to a channel
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		// consume a channel, add the 1 being pushed
		for n := range c {
			sum += n
		}
		// put sum on another channel
		out <- sum
		close(out)
	}()
	return out
}
