package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	defer f.Close()

	if err != nil {
		log.Panic(err)
		return
	}

	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Println("Seed:", seed)
}
