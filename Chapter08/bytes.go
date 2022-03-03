package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("This is"))
	fmt.Fprintf(&buffer, " a string!\n")
	buffer.WriteTo(os.Stdout)
	buffer.WriteTo(os.Stderr) // nothing to print because the buffer variable will be empty afeter the first buffer.writeTo() call

	buffer.Reset()
	buffer.Write([]byte("Mastering Go!"))
	r := bytes.NewReader([]byte(buffer.String()))
	fmt.Println(buffer.String())
	for {
		b := make([]byte, 3)
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}
}
