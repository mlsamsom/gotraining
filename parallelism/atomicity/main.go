package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

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
		// Atomic operations are accessible by one thread at a time
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "counter:", counter)
	}
	wg.Done()
}
