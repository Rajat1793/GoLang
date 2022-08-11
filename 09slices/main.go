package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	var fruitlist = []string{"apple", "mango", "peach"}
	fmt.Printf("Type of fruit list is %T \n", fruitlist)
	fruitlist = append(fruitlist, "banana", "pineapple")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highscore := make([]int, 4)
	highscore[0] = 123
	highscore[1] = 234
	highscore[2] = 345
	fmt.Println(highscore)

	highscore = append(highscore, 111, 222, 444) //it doesnt matter if the size of slice is fixed append will reinitialize the value
	fmt.Println(highscore)

	sort.Ints(highscore) //sorts the sclices
	fmt.Println(highscore)

	//remove a value from slices based on index

	var courses = []string{"react", "go", "javascript", "python", "rust", "r"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
