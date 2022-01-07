/*
Multiple Return Values.
Unlike other languages, Go does not have Exceptions instead error are used.
*/

package main

import (
	"errors"
	"fmt"
	"math"
)

// func returns 2 values, sqrt and error
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers.")
	}

	return math.Sqrt(x), nil
}

func main() {
	// we get 2 values
	result, err := sqrt(16)

	// check if error exist
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
