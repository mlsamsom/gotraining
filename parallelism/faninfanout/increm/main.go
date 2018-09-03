package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := incrementor(2)

	var cnt int
	for n := range c {
		cnt++
		fmt.Println(n)
	}
	fmt.Println("Final Counter:", cnt)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	inc := func(i int) {
		for j := 0; j < 20; j++ {
			c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing:", j)
		}
		done <- true
	}

	// run n goroutines
	for i := 0; i < n; i++ {
		go inc(i)
	}

	// count the finished channels
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()
	return c
}

/*
var count int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementor("1")
	go incrementor("2")
	wg.Wait()
	fmt.Println("Final Counter:", count)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Println("Process: "+s+" printing:", i)
	}
	wg.Done()
}

CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
