// Write a bubble sort in go
package main

import "fmt"

var bubbleSort = func(usList []int) []int {
	sorted := usList
	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	return sorted
}

func main() {
	// Given list
	unSorted := []int{2, 3, 6, 1, 9, 4, 5, 7, 8, 0}
	fmt.Println("Un-sorted List: ", unSorted)
	sortedList := bubbleSort(unSorted)
	fmt.Println("Sorted List: ", sortedList)
}
