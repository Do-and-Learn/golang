package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var root *Node = nil

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{Value: v, Next: nil}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{Value: v, Next: nil}
		return -2
	}
	return addNode(t.Next, v) // t.Next != nil
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
	} else {
		for t != nil {
			fmt.Printf("%d -> ", t.Value)
			t = t.Next
		}
		fmt.Println()
	}
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookupNode(t.Next, v) // t.Next != nil
}

func size(t *Node) int {
	if t == nil {
		return 0
	}
	size := 0
	for t != nil {
		t = t.Next
		size++
	}
	return size
}

func main() {
	fmt.Println(root)
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)
	addNode(root, 100)
	traverse(root)
	fmt.Println(lookupNode(root, 100))
	fmt.Println(lookupNode(root, -100))
}
