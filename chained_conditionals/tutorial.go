package main

import "fmt"

/*
	!
	|| => true || false
	&& => true && true = true
*/

func main() {

	//val := 5 < 7 || 7 > 9
	val := true || false && true // && > ||

	// break it down if it goes crazy :)
	val2 := val || false

	fmt.Printf("%t \n", val2)

}
