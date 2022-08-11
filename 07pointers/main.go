package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// println("value of pointer: ", ptr)
	// fmt.Printf("Variable is of type: %T \n", ptr)

	num := 25
	var ptr = &num
	fmt.Println("address of pointer: ", &ptr)
	fmt.Println("value inside pointer: ", *ptr)
	fmt.Println("value inside pointer: ", ptr)

	*ptr = *ptr + 2
	fmt.Println("value inside pointer: ", *ptr)

}
