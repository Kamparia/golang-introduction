/*
Range: iterate over each element of an array.
*/

package main

import "fmt"

func main() {
	// iterate over an array
	str_arr := []string{"a", "b", "c"}
	for index, value := range str_arr {
		fmt.Println("index:", index, "value:", value)
	}

	// iterate over a map
	vertices := map[string]int{"triangle": 2, "square": 3, "dedecagon": 12 } 
	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}

	// just iterate over the keys
	for key := range vertices {
		fmt.Println("key:", key)
	}	
}