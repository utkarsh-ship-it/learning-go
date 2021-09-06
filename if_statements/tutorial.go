package main

import "fmt"

func main() {

	/*
		if condition {

		}
	*/

	age := 13
	fmt.Println("Before if")
	if age >= 18 {
		fmt.Println("You can ride alone")
	} else if age >= 14 && age < 18 {
		fmt.Println("You can ride with a parent!")
	} else {
		fmt.Printf("Wait %d years to ride with parent or %d years to ride alone. \n", 14-age, 18-age)
	}
}
