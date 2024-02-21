package main

import (
	"fmt"
	"strings"
)

func isPalindrome_(s string) bool {
	// Edge Case Handling
	if len(s) <= 1 {
		return true // Empty string or single-character string is considered a palindrome
	}

	// Preprocess the string: lowercase and remove non-alphanumeric characters
	s = strings.ToLower(s)
	s = removeNonAlphaNumeric(s)

	// Check if the string is a palindrome
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false // Characters at symmetric positions don't match, so it's not a palindrome
		}
		left++
		right--
	}
	return true // All characters match, so it's a palindrome
}

func isPalindrome(s string) bool {
	// Edge Case Handling
	if len(s) <= 1 {
		return true // Empty string or single-character string is considered a palindrome
	}

	// Preprocess the string: lowercase and remove non-alphanumeric characters
	s = strings.ToLower(s)
	s = removeNonAlphaNumeric(s)

	// Create a map to count the occurrences of each character
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	// Check if the string is a palindrome
	oddCount := 0
	for _, count := range charCount {
		if count%2 != 0 {
			oddCount++
		}
		if oddCount > 1 {
			return false // More than one character has an odd count, not a palindrome
		}
	}

	return true // If there is at most one character with odd count, it's a palindrome
}

// Helper function to remove non-alphanumeric characters from the string
func removeNonAlphaNumeric(s string) string {
	var sb strings.Builder
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}

func main() {
	testCases := []struct {
		input  string
		output bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"abba", true},
		{"abca", false},
		{"abcba", true},
		{"ab", true},
		{"A man, a plan, a canal: Panama", true},
	}

	for _, test := range testCases {
		result := isPalindrome(test.input)
		if result != test.output {
			fmt.Printf("Test Failed: %v inputted, %v expected, recieved: %v\n", test.input, test.output, result)
			return
		}
		fmt.Printf("Test Passed: %v inputted, %v expected, recieved: %v\n", test.input, test.output, result)
	}
}
