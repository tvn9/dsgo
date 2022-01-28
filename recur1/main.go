package main

import "fmt"

// run is a recursion fuction
func run(n int) {
	if n < 1 {
		return
	} else {
		println(n)
		run(n - 1)
		fmt.Printf("After recursion call: #%d > %d\n", n, n)
	}
}

func println(n int) {
	fmt.Printf("Before recursion call: #%d > %d\n", n, n)
}

// program starting point
func main() {
	run(100)
}
