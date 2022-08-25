// Find repeated letter in a string and assign how many times a letter is repeated
package main

import "fmt"

func main() {
	// Given string
	fullName := "Thanh Vu Nguyen"

	// Create an emply map
	strMap := make(map[string]int)

	// Insert letter from string to empty map
	// counter := 1
	for _, l := range fullName {
		strMap[string(l)]++
	}

	// Print out map key value pair
	for key, val := range strMap {
		fmt.Println(key, val)
	}
}
