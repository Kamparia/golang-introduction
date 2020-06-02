/*
Functions: 
*/

package main

import "fmt"

// sum function takes 2 integers and return 1 integer
func sum(x int, y int) int {
	return x + y
}

func main() {
	// call external function
	result := sum(2, 3)
	fmt.Println(result)
}