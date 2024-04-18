package main

import "fmt"

type Shaper interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func main() {
	shape1 := Rectangle{height: 1, width: 1}
	shape2 := Circle{radius: 1}

	fmt.Println(shape1)
	fmt.Println(shape2)
}
