package main

import "fmt"

// List struct, to hold head and tail node values
type List struct {
	head *Node
	tail *Node
}

// Node struct, to an int and next Node value
type Node struct {
	value int
	next  *Node
}

// First returns the fist value in List
func (l *List) First() *Node {
	return l.head
}

// Push assigning list with head and tail values
func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &List{}
	for i := 0; i <= 100; i++ {
		l.Push(i)
	}

	n := l.First()
	for {
		fmt.Println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}
}
