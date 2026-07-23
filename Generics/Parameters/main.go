package main 

import (
	"fmt"
	// "sync"
)

func Print[T any](value T){
	fmt.Println("The value of the parameter is : " , value)
}
func main(){
	Print(10)
	Print("100")
	Print(2000)
}