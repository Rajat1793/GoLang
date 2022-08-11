package main

import "fmt"

func main() {
	fmt.Println("Welcome to structure")

	// no inheritance, no child parent in go lang
	rajat := User{"rajat", "rajat@gmail.com", true, 21}
	fmt.Println(rajat)
	fmt.Printf("Details are :%+v\n", rajat)
	fmt.Printf("Name is %v and email is %v\n", rajat.Name, rajat.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
