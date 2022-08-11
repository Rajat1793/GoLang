package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to If/else")

	logincount := 12

	if logincount < 10 {
		fmt.Println("Login is less than 10")
	} else if logincount > 10 {
		fmt.Println("login is greater than 10")
	} else {
		fmt.Println("login is exactly 10")
	}

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num == 3 {
		fmt.Println("Number is 3")
	} else {
		fmt.Println("number is not 3")
	}

}
