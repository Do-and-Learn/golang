package main

import "fmt"

type square struct {
	X float64
}

type circle struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case square:
		fmt.Println("This is a square!")
	case rectangle:
		fmt.Printf("%v is a circle!\n", v) // %v to acquire the value of the type
	case circle:
		fmt.Println("This a a rectangle!")
	default:
		fmt.Printf("Unknown type %T!\n", v)
	}
}

func main() {
	tellInterface(circle{R: 10})
	tellInterface(square{X: 4})
	tellInterface(rectangle{X: 4, Y: 1})
	tellInterface(10)
}
