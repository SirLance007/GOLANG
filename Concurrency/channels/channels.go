// main receiving , worker sending

// package main

// import "fmt"

// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		fmt.Println("Step 3: Goroutine started")
// 		ch <- "hello"
// 		fmt.Println("Step 6: Goroutine finished sending")
// 	}()

// 	fmt.Println("Step 1: Main waiting to receive")

// 	msg := <-ch

// 	fmt.Println("Step 5: Main received:", msg)
// }


// worker sending , main receiving
// package main

// import "fmt"

// func main() {

// 	ch := make(chan string)

// 	go func(){
// 		println("Worker node active")

// 		ch <- "Hello Baby!!"

// 		println("Sended the data succesfully")
// 	}()

// 	println("Recieved the data sucessfully!!")

// 	msg := <-ch

// 	fmt.Printf("Data : %s\n" , msg)
// }

// Buffer
package main 

import "fmt"

func main(){
	ch := make(chan string , 2)

	fmt.Println("Main: Seding A")
	ch <- "A"

	fmt.Println("Main: Sedning B")
	ch <- "B"

	go func(){
		fmt.Println("Worker Started")

		fmt.Println(<-ch)

		fmt.Println(<-ch)
	}()

	fmt.Scanln()
}
