package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func f1(t int) {
	c1 := context.Background() // initialize an empty Context parameter

	// create Done channel (c1.Done) that can be closed.
	// 1. when cancel() is called, or
	// 2. the Done channel of the parent context is closed.
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1():", c1.Err())
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
}

func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
}

func main() {
	// $ go run 16_simpleContext.go 4
	// f1(): 2022-03-10 22:37:53.291948 +0800 CST m=+4.001267251
	// f2(): 2022-03-10 22:37:57.297816 +0800 CST m=+8.007275751
	// f3(): 2022-03-10 22:38:01.298851 +0800 CST m=+12.008450417

	// $ go run 16_simpleContext.go 10
	// f1(): context canceled
	// f2(): context canceled
	// f3(): context canceled
	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	f1(delay)
	f2(delay)
	f3(delay)
}
