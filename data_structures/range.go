/*
	_ is anonymous variable
*/

package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 4, 56, 7, 12, 4, 6, 6}

	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// for i, elements := range a {
	// 	fmt.Println(i, elements)
	// }

	// find duplicates
	// for i, elements := range a {
	// 	for j, elements2 := range a {
	// 		if elements == elements2 && j > i {
	// 			fmt.Println(a[i])
	// 		}
	// 	}
	// }
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}
