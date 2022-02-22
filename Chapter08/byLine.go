package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	var err error

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	err := lineByLine("miktex-console.log")
	if err != nil {
		fmt.Println(err)
	}
}
