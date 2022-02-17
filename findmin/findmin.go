package main

import "fmt"

func main() {
	num := []int{1, 3, 5, 9, 39, 49, 83, 93, 2, 9}
	fmt.Println("slice num:", num)

	min := num[0]
	for _, n := range num {
		if n < min {
			min = n
		}
	}
	fmt.Println("The smallest number in slice num is", min)
}
