package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	lang := make(map[string]string)
	lang["js"] = "javascript"
	lang["r"] = "r"
	lang["py"] = "python"

	fmt.Println("list of language: ", lang)
	fmt.Println("Specifice value: ", lang["r"])

	delete(lang, "r")
	fmt.Println("list of language: ", lang)

	// loop in maps
	for key, value := range lang {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
