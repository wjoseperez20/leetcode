package main

import (
	"fmt"
	"time"
)

func strStr_(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] != needle[0] {
			continue
		}
		if len(haystack[i:]) < len(needle) {
			return -1
		}
		if equal(haystack[i:i+len(needle)], needle) {
			return i
		}
	}
	return -1
}

func equal(s string, expected string) bool {
	if len(s) != len(expected) {
		return false
	}
	for i, v := range []byte(expected) {
		if s[i] != v {
			return false
		}
	}
	return true
}

func strStr(haystack string, needle string) int {
	result := -1

	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return result
	}

	nHystck := len(haystack) - 1
	nNdle := len(needle) - 1
	ptrNdle := 0

	for idx := 0; idx <= nHystck; idx++ {
		if haystack[idx] == needle[ptrNdle] {
			if ptrNdle == nNdle {
				return idx - nNdle
			}
			ptrNdle++
		} else {
			idx -= ptrNdle
			ptrNdle = 0
		}
	}

	return result
}

func main() {
	testCases := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{"Test Case 1", "hello", "ll", 2},
		{"Test Case 2", "aaaaa", "bba", -1},
		{"Test Case 3", "", "", 0},
		{"Test Case 4", "mississippi", "issip", 4},
		{"Test Case 5", "pineapple", "apple", 4},
		{"Test Case 6", "banana", "na", 2},
		{"Test Case 7", "cherry", "rr", 3},
		{"Test Case 8", "blueberry", "blue", 0},
	}

	start := time.Now()
	for _, value := range testCases {
		result := strStr(value.haystack, value.needle)
		fmt.Printf("The starting index of '%s' in '%s' is: %d\n", value.needle, value.haystack, result)
	}
	elapsed := time.Since(start)
	fmt.Printf("strStr function took %s\n", elapsed)

	start = time.Now()
	for _, value := range testCases {
		result := strStr_(value.haystack, value.needle)
		fmt.Printf("The starting index of '%s' in '%s' is: %d\n", value.needle, value.haystack, result)
	}
	elapsed = time.Since(start)
	fmt.Printf("strStr_ function took %s\n", elapsed)
}
