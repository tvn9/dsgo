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
	prev  *Node
}

// First returns the fist value in List
func (l *List) First() *Node {
	return l.head
}

// Last returns the List tail value
func (l *List) Last() *Node {
	return l.tail
}

// Push assigning list with head and tail values
func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

// Next returns the next Node value
func (n *Node) Next() *Node {
	return n.next
}

// Previous
func (n *Node) Previous() *Node {
	return n.prev
}

// The app entry point
func main() {
	l := &List{}
	for i := 0; i <= 10; i++ {
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
	n = l.Last()
	for {
		fmt.Println(n.value)
		n = n.Previous()
		if n == nil {
			break
		}
	}
}
