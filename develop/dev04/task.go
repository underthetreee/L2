package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(words)
	fmt.Println(anagrams)
}

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, word := range words {
		// Convert word to lower case
		wordLower := strings.ToLower(word)
		// sort letters of word
		sortedWord := sortString(wordLower)

		// Add word to anagram set
		anagrams[sortedWord] = append(anagrams[sortedWord], wordLower)
	}

	// Remove sets with one element
	for key, value := range anagrams {
		if len(value) <= 1 {
			delete(anagrams, key)
		}
	}

	return anagrams
}

// Sorts letters of string
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
