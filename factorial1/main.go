package main

import "fmt"

// factorial
func factorial(n int) int {
	if n < 1 {
		return 1
	}
	return n * factorial(n-1)
}

// program starting point
func main() {
	a := factorial(5)
	fmt.Println("Factorial of 5 is:", a)
}
