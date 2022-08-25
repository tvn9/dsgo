// Find repeated letter in a string and assign how many times a letter is repeated
package main

import "fmt"

// find max repeat chars in map
func maxRepChar(charmap map[string]int) map[string]int {
	newMap := make(map[string]int)
	max := 0
	for _, v := range charmap {
		if max < v {
			max = v
		}
	}

	for k, v := range charmap {
		if v == max {
			newMap[k] = v
		}
	}
	return newMap
}

func main() {
	// Given string
	fullName := "Thannnh Vu Nguyen"

	// Create an emply map
	strMap := make(map[string]int)
	// Insert letter from string to empty map
	// counter := 1
	for _, l := range fullName {
		strMap[string(l)]++
	}

	// Print out map key value pair
	for key, val := range strMap {
		fmt.Println("main", key, val)
	}

	// find max repeat char
	newMap := maxRepChar(strMap)
	for k, v := range newMap {
		fmt.Println("The max repeat character is:", k, v)
	}
}
