package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func workdByWord(file string) error {
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
			fmt.Printf("error reading file %s", err)
			return err
		}
		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}
	return nil
}

func main() {
	err := workdByWord("miktex-console.log")
	if err != nil {
		fmt.Println(err)
	}
}
