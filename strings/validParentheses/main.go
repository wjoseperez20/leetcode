package main

import "fmt"

// Best Solution
func isValid(s string) bool {
	stack := make([]rune, 0)

	// Define a mapping of opening and closing brackets
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// Solution Wilmer Perez V2
func isValidV2(str string) bool {
	//Given
	dictionary := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := make(map[rune]int)

	if len(str) <= 1 {
		return false
	}

	for _, value := range str {
		switch value {
		case '(', '[', '{':
			stack[value]++
		case ')', ']', '}':
			if key, flag := dictionary[value]; flag {
				delete(stack, key)
			}
		}
	}

	return len(stack) == 0
}

// Solution WilmerPerez
func isValid_(s string) bool {
	dictionary := map[byte]byte{'(': ')', '[': ']', '{': '}'}

	if len(s) <= 1 {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if dictionary[s[i]] == s[i+1] {
			i++
		} else {
			return false
		}
	}

	return true
}

func main() {
	testCases := []string{"([{}])", "{[())}", "([)]", "{[]}", "{[()]}", "{[(])}", "", "(", "()", "()[]{}", "(])"}

	for _, value := range testCases {
		fmt.Printf("Validation string %v is %v \n", value, isValid(value))
	}

}
