package main

import "fmt"

func smallestSubArrayWithGivenSum_(nums []int, s int) int {
	result := len(nums) + 1
	windowStart, windowSum := 0, 0

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		windowSum += nums[windowEnd]

		for windowSum >= s {
			currentLength := windowEnd - windowStart + 1

			if currentLength < result {
				result = currentLength
			}

			windowSum -= nums[windowStart]
			windowStart++
		}
	}

	if result == len(nums)+1 {
		return 0 // If no subarray found, return 0
	}

	return result
}

func smallestSubArrayWithGivenSum(nums []int, s int) int {
	var result int

	windowSize := 1
	for windowSize < len(nums) {

		for i := 0; i < len(nums)-windowSize+1; i++ {

			result = 1
			sum := 0
			for j := i; j < i+windowSize; j++ {
				sum += nums[j]

				if sum >= s {
					return result
				}

				result++
			}

		}

		windowSize++
		result = 0
	}

	return result
}

func main() {
	testCases := []struct {
		input    []int
		s        int
		expected int
	}{
		{[]int{2, 1, 5, 2, 3, 2}, 7, 2},
		{[]int{1, 2, 3, 4, 5}, 15, 5},
	}

	for _, tc := range testCases {
		actual := smallestSubArrayWithGivenSum_(tc.input, tc.s)
		if actual != tc.expected {
			fmt.Println("Expected", tc.expected, "but got", actual)
		}
	}
}
