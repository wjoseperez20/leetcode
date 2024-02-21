package main

import "fmt"

func longestCommonPrefix_(strs []string) string {
	p := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {
		}
		p = p[:i]
	}
	return p
}

func longestCommonPrefix(strs []string) string {
	// Given
	var result string
	boolFlag := true
	var p1, p2 int

	// First Condition
	if len(strs) == 1 {

		if len(strs[0]) >= 1 {
			return strs[0]
		}

		return result
	}

	for boolFlag {

		// Condition
		if p1 >= len(strs)-1 {
			result += string(strs[p1][p2])
			p2++
			p1 = 0
		}

		if p2 < len(strs[p1]) && p2 < len(strs[p1+1]) {

			if strs[p1][p2] == strs[p1+1][p2] {
				p1++
			} else {
				boolFlag = false
			}

		} else {
			boolFlag = false
		}
	}

	return result
}

func main() {
	testCases := [][]string{
		{"a", "ab"},
		{"flower", "flo", "floght"},
		{""},
		{"a"},
		{"a", ""},
		{"", "beb"},
		{"a", "ab"},
		{"tt", "tt"},
		{"avion"},
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
	}

	for _, value := range testCases {
		fmt.Printf("For %v the common prefix is %v\n", value, longestCommonPrefix(value))
	}
}
