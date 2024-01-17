package main

import (
	"fmt"
)

// Unit conversion constants
func main() {
	value := 1.0
	fromUnit := "banana"
	toUnit := "cm"

	result, err := convert(value, fromUnit, toUnit)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%v %s is equal to %v %s\n", value, fromUnit, result, toUnit)
	}
}
