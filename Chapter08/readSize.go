package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)

	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Panic(err)
	}
	return buffer[0:n]
}

func main() {
	bufferSize := 1000
	f, _ := os.Open("/bin/ls")
	readData := readSize(f, bufferSize)
	if readData != nil {
		fmt.Print(string(readData))
	}
}
