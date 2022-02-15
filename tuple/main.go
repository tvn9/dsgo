//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
	"os"
)

//gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (square, cube int, err error) {
	square = a * a
	cube = a * a * a
	return square, cube, nil
}

// program entry
func main() {
	var square, cube int
	var err error
	square, cube, err = powerSeries(3)
	if err != nil {
		fmt.Fprintln(os.Stderr)
	}

	fmt.Println("Square", square, "Cube", cube)
}
