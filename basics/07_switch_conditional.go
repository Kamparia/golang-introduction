package main

import "fmt"

func main() {
	i := 2

	switch {
		case i == 1, i == 2:
			fmt.Println("one or two")
		case i == 3:
			fmt.Println("three")
		default:
			fmt.Println("Not one, two or three.")
	}
}