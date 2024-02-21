package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

/*
 * Complete the 'countSentences' function below.
 *
 * The function is expected to return a LONG_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY wordSet
 *  2. STRING_ARRAY sentences
 */

/**
How do we tell if a word is an anagram?
- Sort the string, then compare the sorted string

Dictionary
- Store every word (after sorting) from the sentence string array

**/

func countSentences(wordSet []string, sentences []string) []int64 {
	// Dictionary
	var result []int64
	count := make(map[string][]string)

	// Fill Dictionary
	for _, value := range wordSet {
		sorted := sortString(value)
		count[sorted] = append(count[sorted], value)
	}

	// Check Sentences
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")

		total := 1
		for _, word := range words {
			sorted := sortString(word)
			anagrams := count[sorted]
			total *= len(anagrams)

		}

		result = append(result, int64(total))
	}

	return result
}

func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

func main() {
	testCases := []struct {
		wordSet   []string
		sentences []string
		expected  []int64
	}{
		{[]string{"cat", "act", "tab", "bat"}, []string{"cat the bats", "act the bats", "tab the cats"}, []int64{4, 4, 4}},
	}

	for _, value := range testCases {
		fmt.Printf("testPassed?=%v\n", reflect.DeepEqual(countSentences(value.wordSet, value.sentences), value.expected))
	}
}
