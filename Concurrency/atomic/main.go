package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&counter, 1)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&counter, 1)
		}
	}()

	wg.Wait()

	fmt.Println("Counter:", atomic.LoadInt64(&counter))
}