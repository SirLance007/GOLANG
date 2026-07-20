package main

import "fmt"

func main(){
	// Slices -> contains two values : 
	// 1. -> Length 
	// 2. -> Capacity
	// When we breach the size of the array it doubles the size of the array 

	s1 := make([] int, 3 , 5)
	s1[0] , s1[1] , s1[2] = 10 , 20 , 30
	fmt.Println("The Slice s1 : " , s1)
	fmt.Printf("s1 : %v , len : %d , cap : %d\n" , s1 , len(s1) , cap(s1))

	s1 = append(s1 , 9800)
	s1 = append(s1 , 1900)
	// After this the capacity became 10 because it got doubled
	s1 = append(s1 , 1900)
	fmt.Println("Slice array length and capacity : " , len(s1) , cap(s1))

}