package main

import "fmt"

func main() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered:")
		}

	}()

	fmt.Println("Program started")

	panic("Something went wrong")

	fmt.Println("Program ended")
}