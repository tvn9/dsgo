package main

import "fmt"

// Solution 1
func checkPalindrome(str string) bool {
	var palindrome bool
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		if str[i] == str[j] {
			palindrome = true
			fmt.Println(str[i], str[j])
		} else {
			palindrome = false
			fmt.Println(str[i], str[j])
		}
	}
	return palindrome
}

// Solution 2
func checkPalindrome2(str string) bool {
	var palindrome bool
	newStr := ""
	for _, v := range str {
		newStr = string(v) + newStr
		if str == newStr {
			palindrome = true
		} else {
			palindrome = false
		}
	}
	return palindrome
}

func main() {
	// gevin string
	str := "racecar"
	fmt.Println(checkPalindrome(str))
	fmt.Println(checkPalindrome2(str))
}
