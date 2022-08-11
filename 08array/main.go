package main

import "fmt"

func main() {

	fmt.Println("Welcome to Array")
	// 1st way to delcaring array
	var fruitlist = [5]string{"apple", "orange", "", "peach"}
	fmt.Println("Array contains: ", fruitlist)

	// 2nd way of declaring array
	var fruitlist1 [5]string
	fruitlist1[0] = "mango"
	fruitlist1[2] = "grapes"
	fruitlist1[4] = "pinapple"
	fmt.Println("Array contains: ", fruitlist1)

	// 3rd way of declaring array without specifing length of array
	var arr1 = [...]int{1, 2, 3}
	fmt.Println("Array contains: ", arr1)
}
