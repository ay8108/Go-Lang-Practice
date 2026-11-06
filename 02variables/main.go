package main

import "fmt"

const LogingToken string = "ABCD" //Public Variable

func main() {
	var username string = "Ajeet"
	fmt.Println(username)
	fmt.Printf("Variable of type : %T \n", username)

	var isBoolean bool = true
	fmt.Println(isBoolean)
	fmt.Printf("Variable of type : %T \n", isBoolean)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable of type : %T \n", smallVal)

	var smallFloat float64 = 255.785478587988855
	fmt.Println(smallFloat)
	fmt.Printf("Variable of type : %T \n", smallFloat)

	//Default values and some aliases
	var defaultIntval int
	fmt.Println(defaultIntval)
	fmt.Printf("Variable of type : %T \n", defaultIntval)

	var defaultStringval string
	fmt.Println(defaultStringval)
	fmt.Printf("Variable of type : %T \n", defaultStringval)

	//implicit type
	var website = "hello.com"
	fmt.Println(website)
	fmt.Printf("Variable of type : %T \n", website)

	//No var style
	numberOfUsers := 36
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable of type : %T \n", numberOfUsers)

	fmt.Println(LogingToken)
	fmt.Printf("Variable of type : %T \n", LogingToken)
}
