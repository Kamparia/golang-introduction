/*
Compared to other programming languages that have multiple looping 
constructs such as "for", "while", "foreach" etc. The only 
looping construct that exists in Go is the "for" loop construct.
*/

package main

import "fmt"

func main() {
	// method 1
	fmt.Println("Method 1")
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	
	// method 2
	fmt.Println("Method 2")
	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}
	
	// method 3
	fmt.Println("Method 3")
	for k := 0; k <= 10; k++ {
		// skip even values and print odd values
		if k % 2 == 0 {
			continue
		}
		fmt.Println(k)		
	}	
}