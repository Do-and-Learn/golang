package main

import "fmt"

type twoInts struct {
	X int64
	Y int64
}

func reqularFunction(a, b twoInts) twoInts {
	return twoInts{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a twoInts) method(b twoInts) twoInts {
	return twoInts{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a twoInts) method2(b *twoInts) twoInts {
	return twoInts{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a *twoInts) method3(b twoInts) twoInts {
	return twoInts{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a twoInts) method4(b twoInts) *twoInts {
	return &twoInts{X: a.X + b.X, Y: a.Y + b.Y}
}

func main() {
	i := twoInts{X: 1, Y: 2}
	j := twoInts{X: -5, Y: -2}
	fmt.Println(reqularFunction(i, j)) // {-4, 0}
	fmt.Println(i.method(j))           // {-4, 0}
	fmt.Println(i.method2(&j))         // {-4, 0}
	fmt.Println(i.method3(j))          // {-4, 0}
	fmt.Println(i.method4(j))          // &{-4, 0}
}
