package main

import "fmt"

func isSubstring(str1, str2 string) bool {
	n := len(str2)
	for i := 0; i < n; i++ {
		if str2[i] != str1[0] {
			continue
		}

		if len(str2[i:n]) < len(str1) {
			return false
		}

		if check(str2[i:i+len(str1)], str1) {
			return true
		}
	}

	return false
}

func check(str2, str1 string) bool {
	if len(str2) != len(str1) {
		return false
	}

	for key, value := range []byte(str2) {
		if str1[key] != value {
			return false
		}
	}

	return true
}

func main() {
	testCases := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"abc", "dabcd", true},
		{"abc", "dabdb", false},
		{"test", "testing", true},
		{"123", "12345", true},
		{"xyz", "abc", false},
		{"abc", "a", false},
	}

	for _, value := range testCases {
		result := isSubstring(value.str1, value.str2)
		fmt.Printf("The %v is substring of %v = %v . TestPassed?=%v\n", value.str1, value.str2, result, result == value.expected)
	}
}
