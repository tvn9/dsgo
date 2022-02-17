package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	words := []string{"actually", "just", "quite", "really"}

	// words := []string{}

	search, err := search("just", words)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Found %q\n", search)
}

func search(search string, words []string) (string, error) {
	if len(words) <= 0 {
		return "", errors.New("something went wrong")
	}
	for _, w := range words {
		if w == search {
			return search, nil
		}
	}
	return "", nil
}
