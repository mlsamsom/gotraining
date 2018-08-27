package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var m sync.Mutex

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		{ // Lock the global update operation
			m.Lock() // this makes other threads waiting to access below variables wait
			x := counter
			x++
			counter = x
			fmt.Println(s, i, "counter:", counter)
			m.Unlock()
		} // End lock
	}
	wg.Done()
}
