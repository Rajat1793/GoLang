package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Race conditions")

	wg := &sync.WaitGroup{}
	mute := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3) //instead of adding wg.add again and again we can sepcify total number of threads also
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) { //reference of wg
		fmt.Println("Go routine One")
		mute.Lock()
		score = append(score, 1)
		mute.Unlock()
		wg.Done()
	}(wg, mute)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Go routine two")
		mute.Lock()
		score = append(score, 2)
		mute.Unlock()
		wg.Done()
	}(wg, mute)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Go routine three")
		mute.Lock()
		score = append(score, 3)
		mute.Unlock()
		wg.Done()
	}(wg, mute)

	wg.Wait()
	fmt.Println(score)
}
