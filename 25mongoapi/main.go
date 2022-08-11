package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("Welcome to mongoAPI")
	r := router.Router()
	fmt.Println("Server is getting started...")
	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
