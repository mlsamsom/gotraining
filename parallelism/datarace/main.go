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

// Not thread safe, x is copied from counter and updates it after
// some cycles
func incrementor(s string) {
	for i := 0; i < 20; i++ {
		// start not thread safe
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		// end not thread safe
		fmt.Println(s, i, "counter:", counter)
	}
	wg.Done()
}
