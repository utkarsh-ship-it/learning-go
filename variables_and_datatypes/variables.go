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
	//var number uint16 = 260 (Explicit declaration)
	//fmt.Println("Hello World!", number)
	// var number = 260.98 //(Implicit declaration, imiplicitly guesses what type of the variable)
	// fmt.Println("Hello world!", number)

	// this will give us type of number: use Printf
	// fmt.Printf("%T\n", number)

	//number := 6 (fastest and easiest way to declar a variable)
	//fmt.Printf("%T\n", number)

	//number = "Str" (can't do now, since it intantiated as int)

	// create a variable but don't give it a value
	// all of the types has a default value
	var number uint64
	var bl bool // default type false
	fmt.Println(number, bl)

}
