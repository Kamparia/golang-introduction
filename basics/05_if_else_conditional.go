package main

import "fmt"

func output(a int) {
	if a > 0 {
		fmt.Println("The input value", a, "is positive.")
	} else if a < 0 {
		fmt.Println("The input value", a, "is negative.")
	} else {
		fmt.Println("The input value", a, "is zero.")
	}
}

func main() {
	var value = 25
	output(value)
}
