package main 

import (
	"fmt"
	"sync"
)

func main(){
	var counter int 
	var mu sync.Mutex 
	var wg sync.WaitGroup

	wg.Add(2)

	go func(){
		defer wg.Done()

		for i := 0; i < 10 ; i++ {
			mu.Lock()

			counter++

			mu.Unlock()
		}

	}()

	go func(){
		defer wg.Done()

		for i := 0; i < 10 ; i++ {
			mu.Lock()

			counter++

			mu.Unlock()
		}
		
	}()

	wg.Wait()

	fmt.Println("Counter: " , counter)
}