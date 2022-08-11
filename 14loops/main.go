package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	// days := []string{"sunday", "monday", "tuesday", "wednesday", "thrusday", "friday", "saturday"}
	// fmt.Println(days)

	// regular for loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for loop using range
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for each loop
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	// while loop
	value := 1
	for value < 10 {
		if value == 5 {
			break
		}
		if value == 2 {
			goto label
		}
		fmt.Println("Value is: ", value)
		value++
	}

label:
	fmt.Println("Im outside the loop im label")

}
