package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Fast Response"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch2 <- "Slow Response"
	}()

	// Select waits for whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}

	
}