package main

import "fmt"

// Solution 1
func checkPalindrome(str string) bool {
	var palindrome bool
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		if str[i] == str[j] {
			palindrome = true
		} else {
			palindrome = false
		}
	}
	return palindrome
}

// Solution 2
func checkPalindrome2(str string) bool {
	newStr := ""
	for _, v := range str {
		newStr = string(v) + newStr
	}
	return str == newStr
}

// Solution 3
func checkPalindrome3(str string) bool {
	var result bool
	for i, v := range str {
		if v == rune(str[len(str)-(i+1)]) {
			result = true
		} else {
			result = false
		}
	}
	return result
}

func main() {
	// gevin string
	str := "racecar"
	fmt.Println(checkPalindrome(str))
	fmt.Println(checkPalindrome2(str))
	fmt.Println(checkPalindrome3(str))
}
