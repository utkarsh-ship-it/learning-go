package main

import "fmt"

//custom type
type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float64
	//center *Point //embeded struct
	*Point // we can also write this
}

func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	p1 := &Point{y: 1} //passing as reference
	//p1.x = 8 // in struct, we can directly use... normally, we should derefrence it (*p1).x = 8
	var p2 Point = Point{-5, 7}
	changeX(p1)
	//p1.x = 7
	fmt.Println("p1 values: ", p1)
	fmt.Println("x values: ", p1.x, p2.x)
	fmt.Println("y values: ", p1.y, p2.y)

	//draw circle
	c1 := Circle{4.56, p1}

	fmt.Println(c1.y)
}
