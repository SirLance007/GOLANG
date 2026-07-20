package main

import "fmt"

func main(){
	var arr1 [3]int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30

	arr := [3]string{"Prankur" , "Purvi" , "Nitesh"}

	arr3 := [...]bool{true , false}

	fmt.Println("The values of the array are : " , arr3)

	fmt.Println("The array values are : " , arr)
}