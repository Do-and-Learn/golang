package main

import (
	"fmt"
	"sync"
)

var aMutex sync.Mutex

func main() {
	// $ go run -race 15_noRaceC.go
	var waitGroup sync.WaitGroup
	var i int
	numGR := 10

	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(j int) {
			defer waitGroup.Done()
			aMutex.Lock()
			k[j] = j
			aMutex.Unlock()
		}(i)
	}

	waitGroup.Wait()
	k[2] = 10
	fmt.Printf("k = %#v\n", k)
}
