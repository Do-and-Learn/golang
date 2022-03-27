package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createBuffer(buf *[]byte, count int) {
	*buf = make([]byte, count)
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		intByte := byte(random(0, 100))
		if len(*buf) > count {
			return
		}
		*buf = append(*buf, intByte)
	}
}

func Create(dst string, b, f int) error {
	_, err := os.Stat(dst)
	if err == nil {
		return fmt.Errorf("file %s already exists", dst)
	}
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	buf := make([]byte, 0)
	for {
		createBuffer(&buf, b)
		buf = buf[:b]
		if _, err := destination.Write(buf); err != nil {
			return nil
		}
		if f < 0 {
			break
		}
		f = f - len(buf)
	}
	return err
}

var BUFFERSIZE int
var FILESIZE int

func main() {
	output := "/tmp/randomFile"
	BUFFERSIZE, _ = strconv.Atoi("1")
	FILESIZE, _ = strconv.Atoi("100000")
	err := Create(output, BUFFERSIZE, FILESIZE)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(output)
	if err != nil {
		fmt.Println(err)
	}
}
