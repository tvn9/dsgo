package main

import "fmt"

// Queue struct
type Queue struct {
	items chan int
}

// Queue
func (q *Queue) Enqueue(n int) {
	q.items <- n
}

// Dequeue
func (q *Queue) Dequeue() int {
	return <-q.items
}

// program entry
func main() {
	// Initiate q
	q := Queue{
		items: make(chan int, 16),
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(6)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
