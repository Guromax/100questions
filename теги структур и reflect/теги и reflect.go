package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func main() {
	p := Point{10, 20}
	t := reflect.TypeOf(p)
	field, ok := t.FieldByName("X")
	if ok {
		fmt.Println(field.Tag)             // json:"x"
		fmt.Println(field.Tag.Get("json")) // x
	}
}
