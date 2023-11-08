package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// Convert the string to a slice of words.
	words := strings.Fields(s)

	// Declare a map to hold the word count.
	word_count := make(map[string]int)

	// Count the words. Remember that the nil value for an int is 0,
	// so we don't need to initialize a key-value pair like python.
	for _, word := range words {
		word_count[word] += 1
	}

	return word_count
}

func main() {
	wc.Test(WordCount)
}
