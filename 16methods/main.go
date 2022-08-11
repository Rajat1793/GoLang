package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods")

	rajat := User{"rajat", "rajat@gmail.com", true, 21}
	fmt.Println(rajat)
	fmt.Printf("Details are :%+v\n", rajat)
	fmt.Printf("Name is %v and email is %v\n", rajat.Name, rajat.Email)
	rajat.GetStatus()
	rajat.NewEmail()
	fmt.Println("Email changed ?: ", rajat.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user is active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("new email id is: ", u.Email)
}
