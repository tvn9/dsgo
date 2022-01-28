package main

import "fmt"

// factorial
func factorial(n int) {
	var num int
	if n < 1 {
		num = 1
	}
	num = n * (n-1)
	factorial(n - 1)

	fmt.Println(num)
}

// program starting point
func main() {
	factorial(5)
}
