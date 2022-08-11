package main

import "fmt"

const Number int = 20 //Public

func main() {
	fmt.Println("variables")
	var username string = "rajat"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var smallval uint8 = 253
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallfloat float32 = 253.324234234234
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	// default values and alises
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Default value of a Variable is of type: %T \n", anothervariable)

	var anothervariable2 string
	fmt.Println(anothervariable2)
	fmt.Printf("Default value of a Variable is of type: %T \n", anothervariable2)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	numberOfUsers := 300
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type without specifing any type: %T \n", numberOfUsers)

	fmt.Println(Number)
	fmt.Printf("Variable is of type without specifing any type: %T \n", Number)
}
