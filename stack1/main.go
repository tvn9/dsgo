package main

import "fmt"

type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop
func (s *Stack) Pop() int {
	if len(s.items) < 1 {
		return -1
	}

	lEl, aEl := s.items[len(s.items)-1], s.items[0:len(s.items)-1]
	s.items = aEl
	return lEl
}

// program entry
func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	for i, v := range s.items {
		fmt.Printf("#%d push %d\n", i, v)
	}

	for i := len(s.items); i >= len(s.items); i-- {
		fmt.Printf("#%d pop %d\n", i, s.Pop())
	}
}
