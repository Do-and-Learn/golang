package main

import (
	"container/ring"
	"fmt"
)

func main() {
	myRing := ring.New(11)
	fmt.Println("Emptty ring:", *myRing)

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}
	myRing.Value = 2

	sum := 0
	myRing.Do(func(x interface{}) {
		sum += x.(int)
	})
	fmt.Println("Sum:", sum)

	for i := 0; i < myRing.Len()+2; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	println()
}
