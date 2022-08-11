package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer")

	// defer follow last in first out
	// world , one , two
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	mydefer()
}

// stack
// 0, 1, 2, 3, 4
// hello, 43210, two , one , world

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
