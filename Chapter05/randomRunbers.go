package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 100
	TOTAL := 100
	SEED := time.Now().Unix()

	switch len(os.Args) {
	case 2:
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(os.Args[1])
		MAX = MIN + 100
	case 3:
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(os.Args[1])
		MAX, _ = strconv.Atoi(os.Args[2])
	case 4:
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(os.Args[1])
		MAX, _ = strconv.Atoi(os.Args[2])
		TOTAL, _ = strconv.Atoi(os.Args[3])
	case 5:
		MIN, _ = strconv.Atoi(os.Args[1])
		MAX, _ = strconv.Atoi(os.Args[2])
		TOTAL, _ = strconv.Atoi(os.Args[3])
		SEED, _ = strconv.ParseInt(os.Args[4], 10, 64) // int64
	default:
		fmt.Println("Using default values!")
	}

	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		fmt.Print(random(MIN, MAX))
		fmt.Print(" ")
	}
	fmt.Println()
}
