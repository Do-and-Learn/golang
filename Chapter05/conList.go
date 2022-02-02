package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func printList(l *list.List) {
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	values := list.New()
	one := values.PushBack("One")    // One
	two := values.PushBack("Two")    // One Two
	values.PushFront("Three")        // Three One Two
	values.InsertBefore("Four", one) // Three Four One Two
	values.InsertAfter("Five", two)  // Three Four One Two Five
	values.Remove(two)               // Three Four One Five
	values.Remove(two)
	values.InsertAfter("FiveFive", two) // Three Four One Five
	values.PushBack(values)             // Three Four One Five &{{0xc0000c04e0 0xc000020030 <nil> <nil>} 5}

	printList(values)

	values.Init()
	fmt.Printf("After Init(): %v\n", values) // After Init(): &{{0xc000074480 0xc000074480 <nil> <nil>} 0}

	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i))
	}

	printList(values)
}
