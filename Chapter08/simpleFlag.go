package main

import (
	"flag"
	"fmt"
)

func main() {
	// 1st: boolean command line option named k
	// 2nd: default value is true
	// 3rd: usage string displayed with the usage information of the program
	minuxK := flag.Bool("k", true, "k")

	// 1st: int command line option named O
	// 2nd: default value is 1
	// 3nd: usage string displayed with the usage information of the program
	minusO := flag.Int("O", 1, "O")

	flag.Parse() // call after defining the command-line options thay you want to support

	valueK := *minuxK // obtain the value of option k
	valueO := *minusO // obtain the value of option O
	valueO++

	fmt.Println("-k:", valueK)
	fmt.Println("-O:", valueO)
}
