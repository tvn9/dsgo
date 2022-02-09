package main

import "fmt"

// BTree
type Tree struct {
	node *Node
}

// node
type Node struct {
	value int
	left  *Node
	right *Node
}

// Insert
func (t *Tree) insert(v int) *Tree {
	if t.node == nil {
		t.node = &Node{value: v}
	} else {
		t.node.insert(v)
	}
	return t
}

// insert
func (n *Node) insert(v int) {
	if v < n.value {
		if n.left == nil {
			n.left = &Node{value: v}
		} else {
			n.left.insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: v}
		} else {
			n.right.insert(v)
		}
	}
}

// printNode
func printNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	printNode(n.left)
	printNode(n.right)
}

// exists
func (n *Node) exists(v int) bool {
	if n == nil {
		return false
	}
	if n.value == v {
		return true
	}
	if v <= n.value {
		return n.left.exists(v)
	} else {
		return n.right.exists(v)
	}
}

// program entry
func main() {
	t := &Tree{}
	t.insert(10).insert(8).insert(20).insert(9).insert(0).insert(15).insert(25)
	printNode(t.node)

	fmt.Println(t.node.exists(33))
}
