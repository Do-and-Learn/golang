package main

import "fmt"

const SIZE = 10

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return i % size
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			for t := hash.Table[k]; t != nil; t = t.Next {
				fmt.Printf("%d -> ", t.Value)
			}
			fmt.Println()
		}
	}
}

func lookup(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	for t := hash.Table[index]; t != nil; t = t.Next {
		if t.Value == value {
			return true
		}
	}
	return false
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
	fmt.Println(lookup(hash, 119))
	fmt.Println(lookup(hash, 120))
}
