//  Given a string s, sort it in decreasing order based on the frequency of the characters.
// The frequency of a character is the number of times it appears in the string.

// Hashmap Implementation
// Time Complexity: O(n log n)

package main

import (
	"sort"
	"strings"
)

func frequencySort(s string) string {
	freq := make(map[rune]*int)

	for _, char := range s {
		if _, ok := freq[char]; ok {
			(*freq[char])++
		} else {
			count := 1
			freq[char] = &count
		}
	}

	sortedChars := make([]*rune, 0, len(freq))
	for char := range freq {
		sortedChars = append(sortedChars, &char)
	}
	sort.Slice(sortedChars, func(i, j int) bool {
		return (*freq[*sortedChars[i]]) > (*freq[*sortedChars[j]])
	})

	var result strings.Builder
	for _, charPtr := range sortedChars {
		char := *charPtr
		for i := 0; i < (*freq[char]); i++ {
			result.WriteRune(char)
		}
	}

	return result.String()
}
