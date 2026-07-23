package main 

import (
	"fmt"
	"sync"
	"time"
)

func worker (id int , wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println("Worker" , id , "started")

	time.Sleep(time.Second)

	fmt.Println("Worker" , id , "finished")
}

func main(){

	var wg sync.WaitGroup

	wg.Add(3)

	go worker(1 , &wg)
	go worker(2 , &wg)
	go worker(3 , &wg)

	wg.Wait()

	fmt.Println("All workers completed")
}

// package main 

// import (
// 	"fmt"
// 	"sync"
// )

// func main(){
// 	var counter int 
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go func(){
// 		defer wg.Done()

// 		for i := 0 ; i < 1000; i++{

// 		}
// 	}
// }
