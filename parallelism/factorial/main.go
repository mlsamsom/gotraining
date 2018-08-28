package main

import (
	"fmt"
)

func main() {
	f := factorial(4)
	fmt.Println("total:", f)

	fm := factorialMulti(4)
	for n := range fm {
		fmt.Println(n)
	}
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func factorialMulti(n int) chan int {
	// Thread to loop
	c := make(chan int)
	// in separate thread run accumulation loop
	// this is only useful if you're running many of these
	// no map reduce or anything
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()
	return c
}
