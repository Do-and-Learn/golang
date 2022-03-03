package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(time.Second)

	_, ok := <-c // 如果 11 行沒有 close，那麼此行將發生錯誤 fatal error: all goroutines are asleep - deadlock!
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
