package main

import "fmt"

func main() {
	var myInt interface{} = 123
	k, ok := myInt.(int)
	if ok {
		fmt.Println("Success (int):", k)
	}

	v, ok := myInt.(float64)
	if ok {
		fmt.Println("Success (float64):", v)
	} else {
		fmt.Println("Failed without panicking!")
	}

	i := myInt.(int) // type assetion here
	fmt.Println("No checking:", i)

	j := myInt.(bool) // type assertion here, will panic
	// panic: interface conversion: interface {} is int, not bool
	fmt.Println(j)
}
