package main

/*
	map not keep track of the order
*/

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	fmt.Println(mp["apple"])
	mp["utk"] = 100
	//delete(mp, "apple")

	//check if exists
	val, ok := mp["utk"]

	fmt.Println(val, ok)

	// mp := make(map[string]int)
	fmt.Println(len(mp))

}
