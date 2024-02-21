package main

import (
	"log"
	"reflect"
	"time"
)

// Given an array of positive numbers and a positive number ‘k,’ find the maximum sum of any contiguous subarray of size ‘k’.

// The time complexity of the above algorithm will be O(N).
// The algorithm runs in constant space O(1).
func maxSumOfContiguousSubArray_(arr []int, k int) int {
	maxSum := 0
	windowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]

		if windowEnd >= k-1 {
			maxSum = max(maxSum, windowSum)

			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// The time complexity of the above algorithm will be O(N*K).
// The algorithm runs in constant space O(1).
func maxSumOfContiguousSubArray(arr []int, k int) int {
	var m int

	//Sliding Windows
	for i := 0; i < len(arr)-k+1; i++ {

		sum := 0
		for j := i; j < i+k; j++ {
			sum += arr[j]

			if sum > m {
				m = sum
			}
		}
	}

	return m
}

func main() {
	testCases := []struct {
		input    []int
		k        int
		expected int
	}{
		{[]int{2, 1, 5, 1, 3, 2}, 3, 9},
		{[]int{2, 3, 4, 1, 5}, 2, 7},
	}

	start := time.Now()
	for _, tc := range testCases {
		actual := maxSumOfContiguousSubArray(tc.input, tc.k)
		if !reflect.DeepEqual(actual, tc.expected) {
			log.Fatalf("expected %v, but got %v", tc.expected, actual)
		}
	}
	elapsed := time.Since(start)
	log.Printf("Test cases passed in %s", elapsed)

	start = time.Now()
	for _, tc := range testCases {
		actual := maxSumOfContiguousSubArray_(tc.input, tc.k)
		if !reflect.DeepEqual(actual, tc.expected) {
			log.Fatalf("expected %v, but got %v", tc.expected, actual)
		}
	}
	elapsed = time.Since(start)
	log.Printf("Test cases passed in %s", elapsed)
}
