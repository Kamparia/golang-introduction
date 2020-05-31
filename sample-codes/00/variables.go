package main

import "fmt"

func main() {
	// string varible
	name := "John Doe"
	fmt.Println("Customer:", name)
	
	//int variable
	num_items := 100
	fmt.Println("Number of items:", num_items)
	
	// maths operations - float64
	price1, price2 := 9.99, 5.50 // declare multiple variables of the same data_types
	price1_total := price1 * float64(num_items)
	price2_total := price2 * float64(num_items)
	total_cost := price1_total + price2_total

	// print total_cost
	fmt.Println("Total Cost:", total_cost)
}