// Given a multi-digits number - Reverse the numbers
// Examples
// 15 => 51
// 981 => 189
// 500 => 5
// -15 => -51
// -90 => -9

package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func reverseInt(nums []int) []int {
	numArray := []int{}
	for _, n := range nums {
		if n < 0 {
			num := math.Abs(float64(n))
			numStr := numToString(int(num))

			// Reverse String
			revStr := reverseString(numStr)

			// convert string to number
			newNum := stringToInt(revStr)
			numArray = append(numArray, -newNum)

		} else {
			numStr := numToString(n)

			// reverse string
			revStr := reverseString(numStr)

			// convert reversed string back to number
			num := stringToInt(revStr)
			numArray = append(numArray, num)
		}
	}
	return numArray
}

// Reverse a string
func reverseString(str string) string {
	revStr := ""
	for _, s := range str {
		revStr = string(s) + revStr
	}
	return revStr
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func numToString(num int) string {
	str := strconv.Itoa(num)
	return str
}

// Program entry
func main() {
	nums := []int{15, 891, 500, -15, -90}
	fmt.Println(nums)

	n := reverseInt(nums)
	fmt.Println(n)
}
