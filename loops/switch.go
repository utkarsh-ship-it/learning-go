package main

import "fmt"

func main() {
	/*
		ans := 2

		switch ans {
		case 1, 2:
			fmt.Println("1. one")
			fmt.Println("2. two")
		case 3:
			fmt.Println("three")
		default:
			fmt.Println("not a case")
		}
	*/

	ans := -1

	switch {
	case ans > 0:
		fmt.Printf("Greater than 0")
	case ans < 0:
		fmt.Println("Less than 0")
	default:
		fmt.Println("Zero")
	}
}
