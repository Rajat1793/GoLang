package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-01-2006 Monday"))
	fmt.Println(presentTime.Format("01-01-2006 15:04:05 Monday"))

	createDate := time.Date(2002, time.October, 17, 05, 00, 00, 00, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("02-01-2006 Monday"))
	fmt.Println(createDate.Format("01-01-2006 Monday"))
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
 