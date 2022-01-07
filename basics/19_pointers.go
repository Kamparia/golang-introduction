/*
Pointers: Can be used to reference the location of stored value in memory.
*/

package main

import (
	"fmt"
)

func main() {
	val := 20
	ptr := &val // pointer starts with &

	fmt.Println(ptr) // print address in memory where value is stored
	fmt.Println(*ptr) // print the value stored at the address (dereference)

	// change the pointer value
	*ptr = 10
	fmt.Println(*ptr)
	fmt.Println(val)

	// change value itself
	val = 50
	fmt.Println(*ptr)
}