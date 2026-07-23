package main 

import "fmt"

func divide (val1 int , val2 int) (int , error){
	if val2 == 0 {
		return 0 , fmt.Errorf("Cannot divide by zero!!")
	}

	return val1/val2 , nil
}

func main(){
	result , err := divide(10 , 2)
	if err != nil {
		fmt.Println("Error" , err)
	}else{
		fmt.Println("Result : " , result)
	}
}
