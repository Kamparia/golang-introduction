package main

import "fmt"

func main() {
	i := 25

	if i < 50 {

		fmt.Println(i, "is less than 50.")
	
	} else if i > 50 {

		fmt.Println(i, "is greater than 50.")

	} else {

		fmt.Println(i, "is equal to 50.")
		
	}
}