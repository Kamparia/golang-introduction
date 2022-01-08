/*
WaitGroup is used to wait for a group of goroutines to finish.

The below program starts multiple goroutines at the same time and waits for them to finish.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup // create wait group

	for i := 1; i <= 5; i++ {
		wg.Add(1) // add goroutine to wait group
		i := i
		go func() {
			defer wg.Done() // remove goroutine from wait group
			worker(i)
		}()
	}

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines finished.")
}
