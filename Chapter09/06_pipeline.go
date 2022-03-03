package main

import (
	"fmt"
	"math/rand"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			break
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		ok := DATA[x] // 重複出現的數值不會被送至 channel，最後一個數值不會被加總。
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int = 0
	for x := range in {
		sum += x
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	min, max := 10, 20
	rand.Seed(time.Now().UnixNano())

	A := make(chan int)
	B := make(chan int)
	go first(min, max, A)
	go second(B, A)
	go third(B)
	time.Sleep(time.Second)
}
