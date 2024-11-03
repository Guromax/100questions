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
	mycircle := Circle{10}
	mysquare := Square{10}

	circleResult := Shape{mycircle}
	squareResult := Shape{mysquare}

	fmt.Println(circleResult.Area(), squareResult.Area())

}

// ПОЧЕМУ ПО ОТДЕЛЬНОСТИ НЕЛЬЗЯ ???
