/*

int, array is immutable
slice, map are mutable

*/

package main

import "fmt"

/*
func main() {
	var x int = 5
	y := x
	y = 7
	fmt.Println(x, y)
}
*/

func main() {
	var x []int = []int{3, 4, 5}
	y := x
	y[0] = 100
	fmt.Println(x, y)
}
