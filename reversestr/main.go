package main

import (
	"fmt"
)

// Reverse a given string
// "Thanh Nguyen"

// Solution 1
func reverseStr1(str string) string {
	newStr := ""
	for _, s := range str {
		newStr = string(s) + newStr
	}
	return newStr
}

func reverseStr2(str string) string {
	newStr := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		newStr[i], newStr[j] = newStr[j], newStr[i]
	}
	return string(newStr)
}

func main() {
	str := "Thanh Nguyen"
	fmt.Println(reverseStr1(str))
	fmt.Println(reverseStr2(str))
}
