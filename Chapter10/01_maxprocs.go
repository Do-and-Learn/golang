package main

import (
	"fmt"
	"runtime"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	// export GOMAXPROCS=4 可以改變輸出的數值成 4
	fmt.Printf("getGOMAXPROCS(): %v\n", getGOMAXPROCS())
}
