/*

	you can use as return type

*/

package main

import "fmt"

type shape interface {
	name() string
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) name() string {
	return "Rectangle"
}

func (c *circle) name() string {
	return "Circle"
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64 {
	return 3.45 * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5, 7}
	shapes := []shape{&c1, &r1}

	for _, shape := range shapes {
		fmt.Println(shape.name(), getArea(shape))
	}

}
