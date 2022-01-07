/*
Methods:
G0 does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
*/

package main

import (
	"fmt"
)

type Rect struct {
	width, height int
}

type User struct {
	FirstName, LastName string
}

// create methods (area & perimeter) with value type as receiver
func (r Rect) area() int {
	return r.width * r.height
}

func (r Rect) perimeter() int {
	return (2 * r.width) + (2 * r.height)
}

// Greeting create method Greeting for User type
func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	// call method
	r := Rect{width: 10, height: 5}
	fmt.Println("Area:", r.area(), "\nPerimeter:", r.perimeter())

	u := User{"Matt", "Aimonetti"}
	fmt.Println(u.Greeting())
}
