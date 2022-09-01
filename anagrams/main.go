// --- Directions
// Check to see if two provided strings are anagrams of eachother.
// One string is an anagram of another if it uses the same characters
// in the same quantity. Only consider characters, not spaces
// or punctuation.  Consider capital letters to be the same as lower case
// --- Examples
//
//	anagrams('rail safety', 'fairy tales') --> True
//	anagrams('RAIL! SAFETY!', 'fairy tales') --> True
//	anagrams('Hi there', 'Bye there') --> False
package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func anagrams(str1, str2 string) bool {
	cleanStr1 := stringClean(str1)
	cleanStr2 := stringClean(str2)
	newStr1, _ := sortString(cleanStr1)
	newStr2, _ := sortString(cleanStr2)
	return newStr1 == newStr2
}

func sortString(str string) (string, error) {
	sortedStr := ""
	if len(str) > 0 {
		lowcaseStr := strings.ToLower(str)
		sliceStr := strings.Split(lowcaseStr, "")
		sort.Strings(sliceStr)
		sortedStr = strings.Join(sliceStr, "")
	} else {
		return "", errors.New("fail to process string")
	}
	return sortedStr, nil
}

func stringClean(str string) string {
	match := regexp.MustCompile(`\b\W*`)
	return match.ReplaceAllString(str, "")
}

func main() {
	giveStr1 := "RAIL! SAFETY!"
	giveStr2 := "fairy tales"

	result := anagrams(giveStr1, giveStr2)

	fmt.Printf("%q and %q are anagrams? = %t\n", giveStr1, giveStr2, result)
}
