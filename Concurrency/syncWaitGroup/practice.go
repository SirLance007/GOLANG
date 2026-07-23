package main 

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		x := 10
		if(x == 10){
			fmt.Println("The goroutine is completed!!")
		}else{
			fmt.Println("The goroutine is not completed!!")
		}
	}()

	// wg.Wait()
	fmt.Println("The whole function executed successfully!!")
}