package main

import "fmt"

// fibonacci
func fibonacci(n int) int {
	if n < 0 {
		panic("Fibonacci only work with positive integer.")
	} else if n == 1 || n == 2 {
		return n - 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

// program starting point
func main() {
	n := fibonacci(9)
	fmt.Println("Fibonacci of 9 is", n)
}
