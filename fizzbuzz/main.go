// Write a program that console logs the numbers from 1 to n, print fizz if a number is multipiles
// three and print buzz if number is multiples of five, and print fizzBuzz if the number is
// multipiles of both three and five.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("please enter a number")
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(args[0], "is not a number")
		return
	}

	if num < 0 {
		fmt.Println("Please enter a positive number")
		return
	}

	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
