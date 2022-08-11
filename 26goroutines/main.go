package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"}
var waitGroup sync.WaitGroup //pointers
var mute sync.Mutex          //pointers

func main() {
	fmt.Println("Welcome to Go Routines")
	// go greeter("hello") //go is used as a keyword here which created a new thread for the execution
	// greeter("World")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		waitGroup.Add(1) //waitgroup add is been used so that we can add more go routines
	}
	waitGroup.Wait() //wait is used as a method which says wait for go routines to return back
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond) //we have used time here so that we can call the thread execution which is fired up
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer waitGroup.Done() // creates a flag that job is done

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOps in endpoint")
	}
	mute.Lock() // mutex is used to lock the memory for the execution
	signals = append(signals, endpoint)
	mute.Unlock()
	fmt.Printf("%d status code for %s \n", result.StatusCode, endpoint)

}
