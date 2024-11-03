package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	SideLength float64
}

func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}

func main() {
	circle := Circle{Radius: 5}
	square := Square{SideLength: 4}

	shapes := []Shape{circle, square}

	for _, shape := range shapes {
		fmt.Printf("Area of %T, %fn", shape, shape.Area())
	}
}
