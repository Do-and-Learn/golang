package main

import (
	"fmt"
	"time"
)

func main() {
	count := 100
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(time.Second) // 提供 goroutinges 充裕的時間已讓其將資料輸出到螢幕上
}
