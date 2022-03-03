package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	s := []byte("Data to write\n")

	// method 1: fmt.Fprint
	f1, err := os.Create("f1.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f1.Close()
	fmt.Fprint(f1, string(s))

	// method 2: File.WriteString
	f2, err := os.Create("f2.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f2.Close()

	n, err := f2.WriteString(string(s))
	fmt.Printf("Wrote %d bytes\n", n)

	// method 3: Writer.WriteString
	f3, err := os.Create("f3.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f3.Close()

	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", n)
	w.Flush()

	// method 4: ioutil.WriteFile
	f4 := "f4.txt"
	err = ioutil.WriteFile(f4, s, 0644) // 不需要先呼叫 os.Create(..)
	if err != nil {
		log.Panic(err)
	}

	// method 5: io.WriteString
	f5, err := os.Create("f5.txt")
	if err != nil {
		log.Panic(nil)
	}
	n, err = io.WriteString(f5, string(s))
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", n)
}
