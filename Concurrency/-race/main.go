package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	// Goroutine 1: Writer
	go func() {
		for i := 0; i < 1000; i++ {
			counter++ // Unsynchronized Write
		}
	}()

	// Goroutine 2: Writer
	go func() {
		for i := 0; i < 1000; i++ {
			counter++ // Unsynchronized Write
		}
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Final Counter:", counter)
}