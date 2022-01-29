package main

import "fmt"

// Queue hold a list of integer in []int
type Queue struct {
	items []int
}

// Enqueue insert elements to queue
func (q *Queue) Enqueue(n int) {
	q.items = append(q.items, n)
}

// Dequeue removes the queue elements starting index 0 to index n
func (q *Queue) Dequeue() int {
	l := len(q.items)
	if l < 1 {
		return -1
	}
	fEl, nEl := q.items[0], q.items[1:]
	q.items = nEl
	return fEl
}

// program entry point
func main() {
	q := Queue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)

	fmt.Println("Enqueuing...")
	// print out all element inserted
	for i, v := range q.items {
		fmt.Printf("#%d enqueue %d\n", i, v)
	}

	fmt.Println("Dequeuing...")
	// remove all elements in queue starting from q.items[0]
	for i := range q.items {
		fmt.Printf("#%d dequeue %d\n", i, q.Dequeue())
	}

	// Testing dequeuing some more elements expected return value of -1
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	if len(q.items) > 0 {
		panic("dequeue did not remove all elements")
	}

	// Just to verify the tack element == 0
	fmt.Println("Verified len of queue =", len(q.items))
}
