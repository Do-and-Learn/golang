package main

import (
	"log"
	"time"
)

func f1(cc chan chan int, f chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i++ {
			sum += i
		}
		c <- sum
	case <-f: // f is a signal channel
	}
}

func main() {
	cc := make(chan chan int)
	times := 4

	for i := 0; i < times; i++ {
		f := make(chan bool)
		go f1(cc, f)
		ch := <-cc
		ch <- i + 1
		for sum := range ch {
			log.Print("Sum(", i+1, ")=", sum)
		}
		time.Sleep(time.Second)
		close(f)
	}
}
