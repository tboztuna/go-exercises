package main

import (
	"fmt"
)

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

func main() {
	square := square{sideLength: 5}
	triangle := triangle{height: 5, base: 3}

	printArea(square)
	printArea(triangle)
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}
func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func printArea(s shape) {
	fmt.Println("Area of the square is:", s.getArea())
}
