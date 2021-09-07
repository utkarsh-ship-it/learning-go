/*

Two pointers
& - address
* - pointer

*/

package main

import "fmt"

func main() {

	x := 7
	fmt.Println(x) //tell me reference for x

	y := &x
	*y = 9
	fmt.Println(x)

}
