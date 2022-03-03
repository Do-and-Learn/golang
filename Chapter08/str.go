package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("test") // Create read-only Reader from a string
	fmt.Println("r length:", r.Len())

	b := make([]byte, 1)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}

	s := strings.NewReader("This is an error!\n")
	fmt.Println("s length:", s.Len())
	n, err := s.WriteTo(os.Stderr)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Wrote %d bytes to os.Stderr\n", n)
}
