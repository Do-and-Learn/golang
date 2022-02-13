package main

import (
	"fmt"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	fmt.Printf("The type of x is %s.\n", reflect.ValueOf(&x).Elem().Type())

	A := a{100, 200.12, "Struct a"}
	r := reflect.ValueOf(&A).Elem()
	iType := r.Type()
	fmt.Printf("A type is %s\n", iType)
	fmt.Printf("The %d fields are: \n", r.NumField())
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: [%s] with type: [%s] and value: [%v]\n", iType.Field(i).Name, r.Field(i).Type(), r.Field(i).Interface())
	}

	fmt.Println()

	B := b{1, 2, "Struct b", -1.2}
	r = reflect.ValueOf(&B).Elem()
	iType = r.Type()
	fmt.Printf("B type is %s\n", iType)
	fmt.Printf("The %d fields are: \n", r.NumField())
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: [%s] with type: [%s] and value: [%v]\n", iType.Field(i).Name, r.Field(i).Type(), r.Field(i).Interface())
	}
}
