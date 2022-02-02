package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t != nil {
		traverse(t.Left)
		fmt.Print(t.Value, " ")
		traverse(t.Right)
	}
}

func create(n int) *Tree {
	rand.Seed(time.Now().Unix())
	var t *Tree
	for i := 0; i < 2*n; i++ {
		t = insert(t, rand.Intn(n*2))
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{Left: nil, Value: v, Right: nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := create(10)
	fmt.Println(tree.Value)
	traverse(tree)
	fmt.Println()
	insert(tree, -10)
	insert(tree, -2)
	traverse(tree)
}
