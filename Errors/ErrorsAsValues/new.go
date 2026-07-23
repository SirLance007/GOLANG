package main 

import (
	"fmt"
	"errors"
)

func login(username string , password int) (string , error){
	if username == "" {
		return "" , errors.New("Username is requiread!!")
	}else if username == "Prankur" {
		return "Rockstar!!" , fmt.Errorf("This boy is not good !!" , username)
	}
	return  "User logged in successfully" , nil 
}

func main(){
	result , error := login("Prankur" , 123456)
	fmt.Println("Result : " , result)
	fmt.Println("Error : " , error)
}