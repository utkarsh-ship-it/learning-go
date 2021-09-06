// not valid variables
// not starts with number
// no slash
// no keyword
// Implicit and explicit idea
// var name_of_var type
package main

import "fmt"

func main() {
	var name string = "Hello Tim"
	fmt.Println("Hello World!", name)
	//var name string
	// name = "Tim"
	// name = "Bill"
	//var number uint8  = 260 (cannot use 260 (untyped int constant) as uint8 value in variable declaration (overflows)compilerNumericOverflow)
	//var number uint16 = 260
	//fmt.Println("Hello World!", number)
}
