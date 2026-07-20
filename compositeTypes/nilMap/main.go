package main 

import "fmt"

func main(){

	var s []int
	fmt.Println("Slice len:" , len(s))

	// To add elements use append 
	s = append(s , 10)
	fmt.Println("Slice after append : " , s)

	// NIL map behaviour 
	var m map[string]int
	fmt.Println("Map len : " , len(m))

	// Read operation on NIL map
	val := m["test"]
	fmt.Println("Reading NIL map values : " , val)
}