package main

import (
	"fmt"
	"os"
)

func powerSeries(a int) (square, cube int, err error) {
	square = a * a
	// cube = a * a * a
	cube = square * a
	return square, cube, nil
}

func main() {
	var square, cube int
	var err error
	square, cube, err = powerSeries(3)
	if err != nil {
		fmt.Fprintln(os.Stderr)
	}

	fmt.Println("Square", square, "Cube", cube)
}
