package main

import (
	"fmt"
	"runtime"
)

// $ env GOOS=linux GOARCH=arm go build 08_xCompile.go
// $ file 08_xCompile
// 08_xCompile: ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), statically linked, Go BuildID=IO8Ut3P7Yz6Nt9VMF9VF/IP7ig_qts2qwu04Ebr6q/i-e6pycFZ_NtohdDn3p6/fxlJSu1NDD3ZYQCIeC00, not stripped
// $ env GOOS=linux GOARCH=386 go build 08_xCompile.go
// $ file 08_xCompile
// 08_xCompile: ELF 32-bit LSB executable, Intel 80386, version 1 (SYSV), statically linked, Go BuildID=LdrF1W0A0s4KxOMsNQbv/NDbalgxqsmsO4hSXKUR9/QVt5L6sDRc7WDm7EYrHs/jFCGJdCoIdcf6ldBHzte, not stripped

// https://go.dev/doc/install/source

func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("with Go version", runtime.Version())
}
