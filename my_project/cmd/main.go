package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func do() int {
	var x int64
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			atomic.AddInt64(&x, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	return int(x)
}

func main() {
	fmt.Println(do())
}
