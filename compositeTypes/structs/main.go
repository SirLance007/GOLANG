package main

import "fmt"

// Struct definition
type Employee struct {
    Name   string
    Age    int
    Salary float64
}

func main() {
    // 1. Initializing using field names
    emp1 := Employee{
        Name:   "Rahul",
        Age:    28,
        Salary: 70000.50,
    }

    // 2. Zero-value initialization
    emp2 := Employee{
		Name: "Pratyush",
		Age: 30,
		Salary: 12000,
	}

    fmt.Println(emp1 == emp2)

    // 3. Pointer to a struct
    // empPtr := &emp1
    // // Go hume automatically dereference karne deta hai (empPtr.Age aur (*empPtr).Age same hain)
    // empPtr.Age = 29 

    // fmt.Printf("Emp1 Updated Age: %d\n", emp1.Age)
}