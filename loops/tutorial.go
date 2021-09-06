package main

import "fmt"

func main() {

	/*
		x := 0 //define the variable
		for x <= 5 { // put the condition
			fmt.Println(x)
			x++ //increment the variable
		}
	*/

	/*
		for x := 0; x <= 5; x++ {
		 	fmt.Println(x)
		 }
	*/

	// break and continue
	x := 0
	for {

		if x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			x++
			continue
		}
		if x == 999 {
			break
		}
		fmt.Println(x)
		x++
	}

}
