/*
Functions: 
*/

package main

import "fmt"

func main() {
	result := sum(2, 3)
	fmt.Println(result)
}

func sum(x int, y int) int {
	return x + y
}