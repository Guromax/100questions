package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) method() {
	fmt.Println(p.X)
}

type onAmmoPower struct {
	On    bool
	Ammo  int
	Power int
}

func (o *onAmmoPower) Shoot() bool {
	if o.On == false || o.Ammo <= 0 {
		return false
	}
	o.Ammo -= 1
	return true
}

func (o *onAmmoPower) RideBike() bool {
	if o.On == false || o.Power <= 0 {
		return false
	}
	o.Power -= 1
	return true
}

func main() {

	var a [2]string

	a[0] = "hi"

	fmt.Println(a[0], &a[0])

	p1 := Point{
		X: 1,
		Y: 2,
	}

	fmt.Println(p1.X, p1.Y)

	p1.X = 2024

	fmt.Println(p1)

	p1.method()

	player1 := onAmmoPower{
		On:    true,
		Ammo:  12,
		Power: 5,
	}

	fmt.Println(player1)

}
