package main

import (
	"fmt"
	"math"
)

type Shape1 interface {
	Area1() float32
}

type Circle1 struct {
	Radius1 float32
}

func (c Circle1) Area1() float32 {
	return math.Pi * c.Radius1 * c.Radius1
}

type Square1 struct {
	SideLength float32
}

func (s Square1) Area1() float32 {
	return s.SideLength * s.SideLength
}

func main() {
	mycircle := Circle1{Radius1: 12}
	mysquare := Square1{SideLength: 15}

	fmt.Println(mycircle.Area1())

	shapes := []Shape1{mycircle, mysquare}

	fmt.Println(shapes)

	for _, shape := range shapes {
		fmt.Println(shape.Area1())
	}

}
