package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// FAN OUT
	// Distribute the sq work across two goroutines that both read from in.
	// one channel per number
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

// put a bunch of numbers on a channel
func gen(nums ...int) chan int {
	fmt.Printf("TYPE OF NUMS %T\n", nums) // just FYI

	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// concurrently square numbers on a channel
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// consum multiple channels and put on one
func merge(cs ...chan int) chan int {
	fmt.Printf("TYPE OF CS: %T\n", cs) // just FYI

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	// NOTE the scoping,
	// c is being passed as a parameter to the go routine
	// c could be accessable by the go routine but this would point
	// all the new gorountines to the previous c
	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done() // need to count refs so the function doesn't exit
		}(c) // very important, want the goroutine to 'own' c
	}

	// this has to be in a goroutine or else it'll block
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
