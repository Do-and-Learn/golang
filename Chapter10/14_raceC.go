package main

import (
	"fmt"
	"sync"
)

func main() {
	// $ go run -race 14_raceC.go
	var waitGroup sync.WaitGroup
	var i int
	numGR := 10

	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			k[i] = i
		}()
	}

	k[2] = 10
	waitGroup.Wait()
	fmt.Printf("k = %#v\n", k)
}
