package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to Maths")

	// var num1 int = 3
	// var num2 float32 = 324.224
	// fmt.Println("the sum is: ", num1+int(num2))  // its wrong auto conversion as float value will be gone

	//random number using the math
	// rand.Seed(time.Now().UnixNano()) //to get more presise seed value based on the time
	// fmt.Println(rand.Intn(5))

	//random from crypto
	myRandom, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("random number: ", myRandom)

}
