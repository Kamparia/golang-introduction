/*
A goroutine is a lightweight thread managed by the Go runtime.
All programs contains at least one goroutine.

Simple goroutine example with a counter.
The program creates multiple goroutines and increments a counter in each of them.
*/
package main

import (
	"fmt"
	"time"
)

func compute(name string, value int) {
	for i := 1; i <= value; i++ {
		time.Sleep(time.Second) // sleep for 1 second
		fmt.Println(name, ":", i)
	}
}

func main() {
	fmt.Println("Goroutine Tutorial!")

	// input number of goroutines to run
	var num int
	fmt.Print("Enter number of goroutines to run: ")
	fmt.Scan(&num)

	// create goroutines that run compute() function concurrently
	fmt.Printf("Creating %d goroutines...\n", num)
	for i := 1; i <= num; i++ {
		go compute(fmt.Sprintf("Goroutine #%d", i), 10)
	}

	// wait 10secs for all goroutines to finish
	time.Sleep(time.Second * 10)
	fmt.Println("Goroutines finished!")
}
