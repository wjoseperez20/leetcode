package main

import (
	"fmt"
	"math"
)

func getAnagram(s string) int {
	// Count occurrences of each digit in the second half
	count := make(map[rune]int)
	for _, digit := range s[len(s)/2:] {
		count[digit]++
	}

	// Subtract occurrences of each digit in the first half
	// from the counts in the second half
	for _, digit := range s[:len(s)/2] {
		count[digit]--
	}

	// Calculate the total number of operations required
	// as the sum of absolute differences between the counts
	operations := 0
	for _, val := range count {
		operations += int(math.Abs(float64(val)))
	}

	// Return half of the total operations since each operation
	// affects two digits (one from the first half and one from the second half)
	return operations / 2
}

func main() {
	testCases := []struct {
		s      string
		result int
	}{
		{"12345", 1},
		{"123123", 0},
	}

	for _, test := range testCases {
		if result := getAnagram(test.s); result != test.result {
			fmt.Println("Test Failed")
		}
	}
}
