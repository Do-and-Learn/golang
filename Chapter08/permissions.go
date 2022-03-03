package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "miktex-console.log"
	info, _ := os.Stat(filename)
	mode := info.Mode()
	fmt.Println(filename, "mode is", mode.String()[1:10])

	filename = "/dev/random"
	info, _ = os.Stat(filename)
	mode = info.Mode()
	fmt.Println(filename, "mode is", mode.String()[1:10])
}
