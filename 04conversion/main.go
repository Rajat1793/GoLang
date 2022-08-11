package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to store ")
	fmt.Println("Please rate the store ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the raiting: ", input)

	// since the input will be in string format we have to convert into the int for further operations
	// trimspace is used to that we can trim the value and remove all the trailing spaces
	numRaiting, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to raiting", numRaiting+1)
	}
}
