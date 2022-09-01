// Double Link List
package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

// prepend adds data to the list
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// Print the list data
func (l *linkedList) printListData() {
	toPrint := l.head
	for l.length > 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("")
}

//
func (l *linkedList) deleteWithValue(value int) {
	prevToDelete := l.head
	for prevToDelete.next.data != value {
		prevToDelete = prevToDelete.next
	}
	prevToDelete.next = prevToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 2}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListData()
	mylist.deleteWithValue(11)
	mylist.printListData()
}
