package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 100

	var waitGroup sync.WaitGroup

	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	waitGroup.Wait()
}
