package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var start, end int

	for i := 0; i < len(s); i++ {
		// Check for odd-length palindrome with s[i] as center
		len1 := expandAroundCenter(s, i, i)
		// Check for even-length palindrome with s[i] and s[i+1] as center
		len2 := expandAroundCenter(s, i, i+1)

		// Take the maximum length palindrome centered at index i
		maxLength := max(len1, len2)

		// Update start and end indices if we found a longer palindrome
		if maxLength > end-start {
			start = i - (maxLength-1)/2
			end = i + maxLength/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		s        string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for _, value := range testCases {
		fmt.Printf("testPassed?=%v", longestPalindrome(value.s) == value.expected)
	}
}
