package main

import (
	"fmt"
)

type shape interface {
	printArea() string
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

	square.printArea()
	triangle.printArea()
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func (s square) printArea() {
	fmt.Println("Area of the square is:", s.getArea())
}

func (t triangle) printArea() {
	fmt.Println("Area of the triangle is:", t.getArea())
}
