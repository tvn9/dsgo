package main

import "fmt"

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

func main() {
	str := "racecar"

	fmt.Println(checkPalindrome(str))

}
