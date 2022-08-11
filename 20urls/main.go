package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=200"

func main() {
	fmt.Println("Welocme to the URLs")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	// individula params
	qparmas := result.Query()
	fmt.Printf("Type of query params: %T\n", qparmas)
	fmt.Println(qparmas["coursename"])

	for _, val := range qparmas {
		fmt.Println("Params is: ", val)
	}

	// if we want to construct a url based on custom input
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=rajat",
	} //we always pass the reference

	madeUrl := partsOfUrl.String()
	fmt.Println(madeUrl)

}
