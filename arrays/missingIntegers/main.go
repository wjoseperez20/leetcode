package main

import (
	"fmt"
	"reflect"
)

func missingIntegers(input []int, k int) []int {
	numMap := make(map[int]bool)

	for _, num := range input {
		numMap[num] = true
	}

	var missing []int
	for i := 1; len(missing) < k; i++ {
		if _, ok := numMap[i]; !ok {
			missing = append(missing, i)
		}
	}

	return missing
}

func main() {
	testCases := []struct {
		input  []int
		k      int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{6, 7, 8}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{6, 7, 8, 9, 10}},
		{[]int{5, 6, 9, 2, 3}, 5, []int{1, 4, 7, 8, 10}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
	}

	for _, testCase := range testCases {
		result := missingIntegers(testCase.input, testCase.k)
		if reflect.DeepEqual(result, testCase.output) {
			fmt.Printf("Expected: %v, got: %v\n", testCase.output, result)
		}
	}
}
