package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	a := s.getArea()
	fmt.Println("Area of shape is: ", a)
}

func main() {
	t := triangle{
		10,
		5,
	}

	s := square{
		10,
	}

	printArea(t)
	printArea(s)
}
