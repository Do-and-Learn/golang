package main

import "fmt"

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

func addNode(t *Node, v int) *Node {
	if t == nil {
		return &Node{Value: v, Previous: nil, Next: nil}
	}
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return t
	} else if t.Next == nil {
		t.Next = &Node{Value: v, Previous: t, Next: nil}
		return t
	} else {
		return addNode(t.Next, v)
	}
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

func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
	} else {
		var tmp *Node
		for t != nil {
			tmp = t
			t = t.Next
		}
		for tmp != nil {
			fmt.Printf("%d -> ", tmp.Value)
			tmp = tmp.Previous
		}
		fmt.Println()
	}
}

func main() {
	root := addNode(nil, 1)
	traverse(root)
	reverse(root)
	addNode(root, 1)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	traverse(root)
	addNode(root, 100)
	traverse(root)
	reverse(root)
}
