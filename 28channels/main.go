package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels")

	myCh := make(chan int, 2) //second value is buffered value( number of values returned by channel)
	wg := &sync.WaitGroup{}

	// myChannel <- 5 //channel passes the value only when someone is listening to it
	// fmt.Println(<-myChannel)

	wg.Add(2)

	// read only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelClose := <-myCh
		fmt.Println(val)
		fmt.Println(isChannelClose)

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) { //<- arrow makes the channel one directional otherwise its both
		myCh <- 5
		close(myCh)

		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
