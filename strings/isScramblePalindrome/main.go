package main

import "fmt"

func isScrambledPalindrome(s string) bool {
	// Count the frequency of each character
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	// Count the number of characters with odd frequency
	oddCount := 0
	for _, count := range charCount {
		if count%2 != 0 {
			oddCount++
		}
	}

	// Check if the string can be rearranged into a palindrome
	if len(s)%2 == 0 {
		// Even-length string should have all characters with even frequency
		return oddCount == 0
	} else {
		// Odd-length string should have exactly one character with odd frequency
		return oddCount == 1
	}
}

func main() {
	// Test cases
	fmt.Println(isScrambledPalindrome("civicc")) // false
	fmt.Println(isScrambledPalindrome("aabbcc")) // true
	fmt.Println(isScrambledPalindrome("helloo")) // false
	fmt.Println(isScrambledPalindrome("civic"))  // true
	fmt.Println(isScrambledPalindrome("ivicc"))  // true
	fmt.Println(isScrambledPalindrome("aab"))    // true
	fmt.Println(isScrambledPalindrome("abcabc")) // false
}
