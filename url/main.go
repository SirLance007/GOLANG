package main 

const url = "https://google.com"

func main(){

	fmt.Println("LCO Web reuqsts")

	response , err := http.Get(url)

	if err != nil {
		fmt.Println("Error occured" , err)
	}else{
		fmt.Println("Succeess")
	}
}