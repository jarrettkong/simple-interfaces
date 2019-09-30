package main

import (
	"fmt"
	"math"
)

func main() {
	t := triangle{height: 5, base: 7}
	s := square{4}

	printArea(t)
	printArea(s)
}

type triangle struct {
	height float64
	base   float64
}

type square struct{ side float64 }
type shape interface{ getArea() float64 }

func (t triangle) getArea() float64 {
	return t.height * t.base / 2
}

func (s square) getArea() float64 {
	return math.Pow(s.side, 2)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
