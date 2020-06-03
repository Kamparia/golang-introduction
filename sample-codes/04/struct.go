/*
Structs: 
Structs in GO are typed collections of fields. 
They’re useful for grouping data together to form records.
*/

package main

import (
	"fmt"
)

// struct
type employee struct {
	id int
	first_name, last_name, email string
}

func main() {
	a := employee{id: 1, first_name: "John", last_name: "Doe", email: "john.doe@yahoo.com"}
	b := employee{ 2, "Jaden", "Smith", "jsmith@gmail.com"}

	fmt.Println(a)
	fmt.Println(b.email)
}