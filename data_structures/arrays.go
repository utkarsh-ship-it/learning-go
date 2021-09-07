/*
	Arrays are fixed length
	var arr [5]int // default 0s
	[0 0 0 0 0]
	 0 1 2 3 4
	len() - length of an array
*/

package main

import "fmt"

func main() {

	/*
		var arr [5]int

		arr[0] = 100
		arr[4] = 900

		fmt.Println(arr)
	*/

	arr := [3]int{4, 5, 6}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)

	arr2D := [2][3]int{{1, 2, 3}, {3, 3, 4}}
	fmt.Println(arr2D[1][2])

}
