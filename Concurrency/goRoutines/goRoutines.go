package main

import (
	"fmt"
	"time"
)

func printMessage(text string){
	for i := 1 ; i <= 3; i++ {
		fmt.Printf("%s: %d\n" , text , i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main(){
	go printMessage("Goroutine")

	printMessage("Main Thread")
}

