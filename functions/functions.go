/*

	We can label rerurn types/values

	learn more about defer

	functions have their own data types
	function without bracket means reference the function

	have func inside func

	pass func as parameter

	function closure
*/

package main

import "fmt"

// func test() {
// 	fmt.Println("hello")
// }
func add(x, y int) (a int, s int) {
	defer fmt.Println("hello!")
	a = x + y
	s = x - y
	return
}

func test2(mfunc func(int) int) {
	fmt.Println(mfunc(7))
}

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	// ans, sub := add(10, 2)
	// fmt.Println(ans, sub)

	// x := test
	// x()

	// test := func(x int) int {
	// 	return x * -1
	// }
	// test3 := func(x int) int {
	// 	return x * 7
	// }
	//test2(test3)
	returnFunc("hello")()
}
