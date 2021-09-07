/*

Two pointers
& - pointer to get address
* - dereference operator

*/

package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	// toChange := "hello"
	// changeValue(&toChange)
	// fmt.Println(toChange)
	toChange := "hello"
	changeValue2(toChange)
	fmt.Println(toChange)

}
