package main 

import (
	"fmt"
	// "time"
)

func producer(ch chan int){
	for i := 1; i<= 3; i++ {
		// ch <- i*10
		// time.Sleep(100*time.Millisecond)
	}
	// Aagar apan channel ban nhi kargee toh ye infinite looop mai phass jayega
	close(ch)
	// for i := 1; i<= 3; i++ {
	// 	ch <- i*10
	// 	// time.Sleep(100*time.Millisecond)
	// }
}

func main(){
	ch := make(chan int)

	go producer(ch)

	// range reads until channel is closed 
	for val := range ch{
		fmt.Println("Received: " , val)
	}

	fmt.Println("Loop finished cleanly!!")
}