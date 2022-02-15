//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import "fmt"

//gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

// program entry
func main() {
	var square, cube int
	square, cube = powerSeries(3)

	fmt.Println("Square", square, "Cube", cube)
}
