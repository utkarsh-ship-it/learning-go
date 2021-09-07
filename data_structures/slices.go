/*
Slices

length and capacity same for arrays

0 [6 3] 9

ptr points to 6
length of slice [6 3] = 2
capacity = 4 // in case 0 6 3 9 12

slice operator :
*/

package main

import "fmt"

func main() {

	var x [5]int = [5]int{1, 2, 3, 4, 5}

	//var s []int = x[:] // copy the array

	var s []int = x[1:3] // not include 3

	fmt.Println(cap(s)) // 4 because we start with index 1 then rest of the elements

	var a []int = []int{5, 6, 7, 8, 9}
	fmt.Println(cap(a[:3]))

	// use of make
	m := make([]int, 5)
	fmt.Printf("%T \n", m)
}
