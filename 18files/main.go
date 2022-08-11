package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to Files")

	writeString()
	readFile("./testfile.txt")

}
func writeString() {
	content := "This need to be in inside the file"

	file, err := os.Create("./testfile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is: ", length)
	defer file.Close()
}

func readFile(filename string) {
	// whenever we read a file it always return the data byte format not the string format
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Lines read from a file:\n", string(databyte))// string is used to conver the databyte into string
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
