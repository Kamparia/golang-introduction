/*
Channels are a way to communicate (data sharing) between goroutines.

Application to share data between goroutines (hello & receive)
Both goroutines send and receive data via 2 channels (done & anotherDone).
*/
package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func receive(receive bool, anotherDone chan bool) {
	if receive == true {
		fmt.Println("from receive")
		anotherDone <- true
	}
}
func main() {
	// create 2 channels of type bool
	done := make(chan bool)
	anotherDone := make(chan bool)

	// create 2 goroutines and pass the channels to them
	go hello(done)
	toSend := <-done
	go receive(toSend, anotherDone)
	<-anotherDone
	fmt.Println("main function")
}
