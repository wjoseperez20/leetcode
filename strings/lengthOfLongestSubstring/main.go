package main

import (
	"fmt"
	"time"
)

func lengthOfLongestSubstring(s string) int {
	var (
		maxLength int
		start     int
		end       int
	)

	for end < len(s) {
		for i := start; i < end; i++ {
			if s[i] == s[end] {
				start = i + 1
				break
			}
		}

		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}

		end++
	}

	return maxLength
}

func lengthOfLongestSubstring_(s string) int {
	charMap := make(map[rune]int)
	start, maxLength := 0, 0

	// Iterate over the string
	for end, char := range s {

		if lastSeen, ok := charMap[char]; ok && lastSeen >= start {
			start = lastSeen + 1
		}

		// If the length of the current substring is greater than maxLength, update maxLength
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}

		// Add the character and its position to the map
		charMap[char] = end
	}

	// Return the length of the longest substring without repeating characters
	return maxLength
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	start := time.Now()
	for _, testCase := range testCases {
		fmt.Printf("input: %v, expected: %v, output: %v\n", testCase.input, testCase.expected, lengthOfLongestSubstring(testCase.input))
	}
	elapsed := time.Since(start)
	fmt.Println("Time taken lengthOfLongestSubstring: ", elapsed)

	start = time.Now()
	for _, testCase := range testCases {
		fmt.Printf("input: %v, expected: %v, output: %v\n", testCase.input, testCase.expected, lengthOfLongestSubstring_(testCase.input))
	}
	elapsed = time.Since(start)
	fmt.Println("Time taken lengthOfLongestSubstring_: ", elapsed)
}
