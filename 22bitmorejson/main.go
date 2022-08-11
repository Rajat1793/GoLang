package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //in structure if we have given the 3 paramater while defining it acts as a alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //we give - after json so that password field is hidden, no other symbol can be used
	Tags     []string `json:"tags,omitempty"` //omitempty is given so that nil value is not reflected
}

func main() {
	fmt.Println("Welcome to Json")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React JS", 299, "lco", "asd", []string{"web-dev", "js"}},
		{"Javascript JS", 199, "lco", "aqeesd", []string{"fullstack", "js"}},
		{"Bootstarp JS", 999, "lco", "asfbd", nil},
	}
	// package this data as json
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	//data returning from the web will be in form of byte and we have to convert it into string or we might have to use it as json data
	jsonDataFromWeb := []byte(` 
	{
		"coursename": "React JS",
		"Price": 299,
		"website": "lco",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb) //json.valid returns the boolean value
	if checkValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) //we are giving & so the reference is passed and not the copy
		fmt.Printf("%#v\n", lcoCourse)              //%#v to print the interface
	} else {
		fmt.Println("JSON is not valid")
	}

	// some cases where we want just to add data to key value

	var myOnlineData map[string]interface{} //interface type is defined as a set of method signature
	// whenever we create a map for online json data  first value will be key value pair and later value can be anything int string
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T\n", k, v, v)
	}

}
