package main

import "fmt"

func main() {
	// name of user
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	// age of user
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	// print user details
	fmt.Printf("Hello %s you are %d years old.\n", name, age)
}
