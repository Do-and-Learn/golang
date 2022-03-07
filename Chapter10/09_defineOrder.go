package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a // is blocked by the channel stored in the a parameter
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b) // unblock an other function
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)

	close(x)
	time.Sleep(3 * time.Second)
}
