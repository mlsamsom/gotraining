package main

import (
	"fmt"
	"sync"
)

func main() {
	dat := gen(100000)

	// Distribute work across 10 channels
	allChans := make([]chan int, 10, 10)
	for cID := 0; cID < 10; cID++ {
		allChans[cID] = factorial(dat)
	}

	var cnt int
	for n := range merge(allChans...) {
		cnt++
		fmt.Println(cnt, "\t", n)
	}
	/*
		c0 := factorial(dat)
		c1 := factorial(dat)
		c2 := factorial(dat)
		c3 := factorial(dat)
		c4 := factorial(dat)
		c5 := factorial(dat)
		c6 := factorial(dat)

		for n := range merge(c0, c1, c2, c3, c4, c5, c6) {
			fmt.Println(n)
		}
	*/

}

// Generator for numbers to test
func gen(howMany int) chan int {
	howMany = howMany / 10
	out := make(chan int)
	go func() {
		for i := 0; i < howMany; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

// Runs the factorial calculation in parrallel on each member of the input channel
func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
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
