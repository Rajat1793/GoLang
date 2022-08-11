package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Mod")
	greeter()
	// below lines are used to handle the route
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	// below command is used to log in case of error
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) { //if we want to read any request its done via r and if anything has to be writtern it is done via w writter
	w.Write([]byte("<h1>Welcome to Go lang</h1>"))
}
