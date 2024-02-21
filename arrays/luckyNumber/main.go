package main

import "fmt"

func luckyNumber(str string) string {
	if len(str) < 1 {
		return str
	}

	s := []rune(str)
	n := len(s)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if s[j] < s[minIdx] {
				minIdx = j
			}
		}
		// Swap the minimum element with the current element
		s[i], s[minIdx] = s[minIdx], s[i]
	}

	return string(s)
}

func main() {
	testCases := map[string]string{
		"231":        "123",
		"642":        "246",
		"54321":      "12345",
		"9876543210": "0123456789",
		"11111":      "11111",
	}

	for key, value := range testCases {
		result := luckyNumber(key)
		fmt.Printf("The lucky number for %v is %v . TestPassed?=%v\n", key, result, result == value)
	}
}
