/*
Maps: holds key-value pairs. Similar to a dictionary in Python.
*/

package main

import "fmt"

func main() {
	//keys are strings while values are int
	vertices := map[string]int{"triangle": 2, "square": 3, "dedecagon": 12 } 
	
	// delete item
	delete(vertices, "square")
	fmt.Println(vertices)

	// get value by key
	fmt.Println(vertices["triangle"])

	// check if key is present in map
	_, is_val_present := vertices["triangle"]
	fmt.Println(is_val_present) // prints true or false
}