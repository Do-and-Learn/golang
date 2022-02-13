package main

import (
	"ch07/myInterface"
	"fmt"
	"math"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

func Calculate(x myInterface.Shape) {
	fmt.Println("Area:", x.Area())
	fmt.Println("Perimeter", x.Perimeter())
}

func main() {
	Calculate(square{X: 10})
	Calculate(circle{R: 5})
}
