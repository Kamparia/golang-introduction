package main

import "fmt"

func main() {
	// string varible
	var name string = "John Doe" //declaration type 1
	fmt.Println("Customer:", name)

	//int variable
	num_items := 100 //declaration type 2
	fmt.Println("Number of items:", num_items)

	// maths operations - float64
	price1, price2 := 9.99, 5.50 //declaration type 3
	price1_total := price1 * float64(num_items)
	price2_total := price2 * float64(num_items)
	total_cost := price1_total + price2_total

	// print total_cost
	fmt.Println("Total Cost:", total_cost)
}
