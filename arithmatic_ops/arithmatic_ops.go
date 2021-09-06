package main

import (
	"fmt"
)

func main() {

	var num1 int = 9
	var num2 int = 4

	//answer := num1 / num2

	//fmt.Printf("%d", answer) // give you integer representation
	//fmt.Printf("%g", answer)

	// Acronym B - Bracket E - Exponents D - Division M - Multiplication A - Addition S - Subtraction (basic order of operations)
	answer := num1 % num2
	fmt.Printf("%d", answer)

	// Checkout math package
}
