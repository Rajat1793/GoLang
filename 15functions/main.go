package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeter()
	// anonomus funtion
	func() {
		fmt.Println("Im anomonus function")
	}()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	result2, message := add(1, 3, 4, 5, 6)
	fmt.Println("Result 2 is: ", result2)
	fmt.Println("Result 2 is: ", message)

}

func greeter() {
	fmt.Println("Im greeter function")
}

// function declaeration its called funtion signatures
func adder(valone int, valtwo int) int {
	return valone + valtwo
}

// funtion with n number of argument
func add(value ...int) (int, string) {
	total := 0
	for _, val := range value {
		total += val
	}
	return total, "Its a adder funtion"
}
