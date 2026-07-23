package main 

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println("Worker Finished!!")
}

func main(){
	var wg sync.WaitGroup 
	wg.Add(1)

	go work(&wg)

	wg.Wait()

	fmt.Println("Main Finished")
}