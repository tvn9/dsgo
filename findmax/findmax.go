package main

import "fmt"

func main() {
	num := []int{1, 3, 5, 9, 39, 49, 83, 93, 2, 9}
	fmt.Println("slice num:", num)

	max := num[0]
	for _, n := range num {
		if n > max {
			max = n
		}
	}
	fmt.Println("The largest number in slice num is", max)
}
