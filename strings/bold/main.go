package main

import (
	"fmt"
	"strings"
	"time"
)

// Best Solution
func embolden(text string, substrings []string) string {
	// TestNoSubstrings check
	if len(substrings) < 1 {
		return text
	}

	// Create a boolean array to mark the positions to be emboldened
	boldPositions := make([]bool, len(text))

	// Iterate through substrings
	for _, sub := range substrings {
		// Find all occurrences of the substring in the text
		startIndex := 0
		for {
			index := strings.Index(text[startIndex:], sub)
			if index == -1 {
				break
			}

			// Mark positions to be emboldened
			for i := startIndex + index; i < startIndex+index+len(sub); i++ {
				boldPositions[i] = true
			}

			startIndex += index + 1
		}
	}

	// Construct the result string with bold tags
	var result strings.Builder
	isBold := false
	for i := 0; i < len(text); i++ {
		if boldPositions[i] {
			if !isBold {
				result.WriteString("<b>")
			}
			isBold = true
		}

		result.WriteByte(text[i])

		if isBold && (i == len(text)-1 || !boldPositions[i+1]) {
			result.WriteString("</b>")
			isBold = false
		}
	}

	return result.String()
}

// Solution Wilmer Perez
func embolden_(text string, substrings []string) string {
	// Given
	var target string
	var result []string

	// TestNoSubstrings check
	if len(substrings) < 1 {
		return text
	}

	// TestEmptyText
	if len(text) < 1 {
		return text
	}

	// Iteration Targets
	k := 0
	for i := 0; i <= len(substrings)-1; i++ {

		target = substrings[i]
		if len(target) <= len(text) {

			j := 0
			for j <= len(target)-1 && k <= len(text)-1 {

				if target[j] == text[k] {
					if j == 0 {
						result = append(result, "<b>"+string(text[k]))
						j++
					} else if j == len(target)-1 {
						result = append(result, string(text[k])+"</b>")
						j = 0
						k++
						break
					} else {
						result = append(result, string(text[k]))
						j++
					}

					k++

				} else {
					result = append(result, string(text[k]))
					k++
				}
			}
		}
	}

	if k < len(text)-1 {
		for index := k; index < len(text); index++ {
			result = append(result, string(text[index]))
		}
	}

	resultString := strings.Join(result, "")

	return resultString
}

func main() {
	startTime := time.Now()
	TestNoSubstrings()
	TestOneSubstring()
	TestMultipleSubstrings()
	TestEmptyText()
	TestLongText()
	elapsedTime := time.Since(startTime)

	fmt.Printf("\nPerformance Test: %s\n", elapsedTime)
}

func TestNoSubstrings() {
	fmt.Print("\nTestNoSubstrings - ")
	var result = embolden("abcxyz", []string{})
	fmt.Print(result == "abcxyz")
}

func TestOneSubstring() {
	fmt.Print("\nTestOneSubstring - ")
	var result = embolden("abcxyz", []string{"abc"})
	fmt.Print(result == "<b>abc</b>xyz")
}

func TestMultipleSubstrings() {
	fmt.Print("\nTestMultipleSubstrings - ")
	var result = embolden("abcxyz123", []string{"abc", "123"})
	fmt.Print(result == "<b>abc</b>xyz<b>123</b>")
}

func TestEmptyText() {
	result := embolden("", []string{"abc", "123"})
	fmt.Print("\nTestEmptyText - ", result == "")
}

func TestLongText() {
	result := embolden("Lorem ipsum dolor sit amet, consectetur adipiscing elit.", []string{"ipsum", "consectetur"})
	fmt.Print("\nTestLongText - ", result == "Lorem <b>ipsum</b> dolor sit amet, <b>consectetur</b> adipiscing elit.")
}
