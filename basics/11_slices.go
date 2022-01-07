/*
Slices:
An array has a fixed size. A slice, on the other hand, is a 
dynamically-sized, flexible view into the elements of an array. 
In practice, slices are much more common than arrays.
*/

package main

import "fmt"

func main() {
	// initialize a slice with 3 elements
	s := []int{0, 0, 0}
	fmt.Println(s)
	
	// index into the array
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println(s)

	// print
	fmt.Println(s[0])
	fmt.Println(len(s))
	fmt.Println(s[1:3])
	fmt.Println(s[:3])

	// append function
	s = append(s, 4, 5)
	fmt.Println(s)
	
	// copy function
	x := make([]int, len(s))
	copy(x, s)
	
	x[0] = 500
	fmt.Println(x)
	
	// multi-dimenesional (2D) slice
	ss := [][]uint8{
	    {0, 1, 2, 3},
	    {4, 5, 6, 7},
	}
	fmt.Println(ss)
}