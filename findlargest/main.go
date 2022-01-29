package main

import "fmt"

// findLargestNum(n ...int)
func findLargestNum(a []int, n int) int {
	var highest = 1
	if a[n] == -1 {
		return highest
	} else {
		if a[n] > highest {
			highest = a[n]
		}
	}
	return findLargestNum(a, n-1)

}

// program entry
func main() {
	a := []int{23, 234, 34, 5, 6, 79, 49, 20, 2, 882, 355}
	n := len(a) - 1
	l := findLargestNum(a, n)
	fmt.Println(l)
}
