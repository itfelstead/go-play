package main

import "fmt"

type triangle struct {
	base   int
	height int
}

type square struct {
	length int
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{5, 5}
	s := square{5}

	//s.printArea() // won't work as can't call method from non-concrete type
	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {
	return float64(s.length * s.length)
}

func (t triangle) getArea() float64 {
	return float64((t.base * t.height) / 2)
}

//func (s shape) printArea() { // won't work as can't use Interface type as receiver
func printArea(s shape) {

	fmt.Println("Area:", s.getArea())
}
