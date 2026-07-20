package main 

import "fmt"

func main(){
	original := []int {1 , 2 , 3 , 4}

	subSlice := original[0:3]
	fmt.Println("Before modification:")
    fmt.Println("Original :", original)
    fmt.Println("SubSlice :", subSlice)

	subSlice[0] = 99
	fmt.Println("Before modification:")
    fmt.Println("Original :", original)
    fmt.Println("SubSlice :", subSlice)
}