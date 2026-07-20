package main

import "fmt"

func main(){
	employeeSalary := make(map[string]int)
	employeeSalary["Rahul"] = 50000
	employeeSalary["Rohit"] = 60000

	scores := map[string]int{
		"TeamA" : 10,
		"TeamB" : 20,
	}

	// Comma-ok idiom (key check karne ke liye)
	salary, exits := employeeSalary["Amit"]
	if exits {
		fmt.Println("Amit's salary is :" , salary)
	}else{
		fmt.Println("Amit does not exit in the map! Default value received:" , salary)
	}

	delete(employeeSalary , "Rohit")
	fmt.Println("After delete Rohit :" , employeeSalary)
	fmt.Println("Score map:" , scores)
}