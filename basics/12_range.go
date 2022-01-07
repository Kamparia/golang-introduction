/*
Range: 
Iterates over each element of a slice or map. When ranging over a slice, 
two values are returned for each iteration. The first is the index, 
and the second is a copy of the element at that index.
*/

package main

import "fmt"

func main() {
	// iterate over a slice
	letters := []string{"a", "b", "c"}
	for index, value := range letters {
		fmt.Println("index:", index, "value:", value)
	}

	// iterate over a map
	vertices := map[string]int{"triangle": 2, "square": 3, "dedecagon": 12 } 
	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}

	// iterate over map (keys only)
	for key := range vertices {
		fmt.Println("key:", key)
	}	
}