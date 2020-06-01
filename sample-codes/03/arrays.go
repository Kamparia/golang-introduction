/*
Arrays: have fixed length.
*/

package main

import "fmt"

func main() {
	// create int array of 5 elements
	int_arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(int_arr)

	// print given element of an array
	fmt.Println(int_arr[1])

	// print lenght of array
	fmt.Println(len(int_arr))

	// change the values of array elements using the index
	int_arr[0] = 5
	int_arr[1] = 10
	fmt.Println(int_arr)

}