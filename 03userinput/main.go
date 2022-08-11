package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to User input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin) //input from the user
	fmt.Println("Enter the raiting: ")
	//comma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating: ", input)
	fmt.Printf("Type of input is: %T", input)
}
