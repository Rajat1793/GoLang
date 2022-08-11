package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web verb request")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content lenght: ", response.ContentLength)

	var responseString strings.Builder //it contains the raw data
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is :", byteCount)
	fmt.Println("String data is: ", responseString.String())

	// to convert bytecode into string
	// fmt.Println(string(content))
}

func PerformPostRequest() {

	const myurl = "http://localhost:1111/post"

	// fake json data
	requestBody := strings.NewReader(`
		{
			"coursename": "Go Lang",
			"price" : 0,
			"platform" : "LCO"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:1111/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "rajat")
	data.Add("lastname", "jaiswal")
	data.Add("email", "rajat@go.dev")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
