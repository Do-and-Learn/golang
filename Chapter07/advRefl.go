package main

import (
	"fmt"
	"os"
	"reflect"
)

type t1 int // go 將此型別表示成 main.t1
type t2 int // go 將此型別表示成 main.t2

type a struct {
	X    int
	Y    float64
	Text string
}

func (a1 a) compareStruct(a2 a) bool {
	r1 := reflect.ValueOf(&a1).Elem()
	r2 := reflect.ValueOf(&a2).Elem()

	for i := 0; i < r1.NumField(); i++ {
		if r1.Field(i).Interface() != r2.Field(i).Interface() {
			return false
		}
	}
	return true
}

func printMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Printf("Type to examine: %s\n", t)

	for j := 0; j < r.NumMethod(); j++ {
		method := t.Method(j)
		fmt.Println(method.Name, "->", method.Type)
	}
}

func main() {
	x1 := t1(100)
	x2 := t2(100)
	fmt.Println("The type of x1 is", reflect.TypeOf(x1)) // main.t1 而非 int
	fmt.Println("The type of x2 is", reflect.TypeOf(x2)) // main.t2 而非 int

	var p struct{}
	r := reflect.New(reflect.ValueOf(&p).Type()).Elem()
	fmt.Println(reflect.ValueOf(&p).Type())            // *struct{}
	fmt.Println("The type of r is", reflect.TypeOf(r)) // reflect.Value

	a1 := a{1, 2.1, "A1"}
	a2 := a{1, -2, "A2"}

	if a1.compareStruct(a1) {
		fmt.Println("Equal!")
	}
	if !a1.compareStruct(a2) {
		fmt.Println("Not Equal!")
	}
	var f *os.File
	printMethods(f)
}
