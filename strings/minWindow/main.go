package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if t == "" || s == "" {
		return ""
	}

	need := make(map[rune]int)
	for _, ch := range t {
		need[ch]++
	}

	have := 0
	minLen := math.MaxInt32
	minStart := 0

	left, right := 0, 0
	window := make(map[rune]int)

	for right < len(s) {
		ch := rune(s[right])
		window[ch]++

		if count, ok := need[ch]; ok && window[ch] == count {
			have++
		}

		for have == len(need) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}

			c := rune(s[left])
			window[c]--

			if count, ok := need[c]; ok && window[c] < count {
				have--
			}

			left++
		}

		right++
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[minStart : minStart+minLen]
}

func main() {
	testCases := []struct {
		input    string
		t        string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
	}

	for _, tc := range testCases {
		actual := minWindow(tc.input, tc.t)
		if actual != tc.expected {
			fmt.Printf("Error: expected %s but got %s\n", tc.expected, actual)
		}
	}
}
