package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)

	// Send operations
	ch <- "Go"
	// ch <- "Golang"

	// Close operation
	// Deleted the channel data
	close(ch)

	// close(ch)

	// Receive operations (with comma-ok idiom)
	// Top of the queue
	val1, ok1 := <-ch
	fmt.Printf("Val: %s, OK: %t\n", val1, ok1) // Data retrieved

	// Top of the queue
	val2, ok2 := <-ch
	fmt.Printf("Val: %s, OK: %t\n", val2, ok2) // Data retrieved

	// Reading after channel is empty AND closed
	val3, ok3 := <-ch
	fmt.Printf("Val: '%s', OK: %t\n", val3, ok3) // Returns zero-value & false
}
