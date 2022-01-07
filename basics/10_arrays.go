/*
Arrays in GO have a fixed length.
*/

package main

import "fmt"

func main() {
	// create int array of 5 integer elements
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// print given element of an array
	fmt.Println(arr[1])

	// print lenght of array
	fmt.Println(len(arr))

	// change the values of array elements using the index
	arr[0] = 5
	arr[1] = 10
	fmt.Println(arr)

	// create a slice of int array
	slice := arr[1:3]
	fmt.Println(slice)
}
