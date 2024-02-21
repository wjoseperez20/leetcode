package main

import (
	"fmt"
	"time"
)

func majorityElement(nums []int) int {
	var result int

	if len(nums) <= 1 {
		return nums[result]
	}

	majority := len(nums) / 2
	storage := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		storage[nums[i]]++

		if storage[nums[i]] > majority {
			result = nums[i]
		}
	}

	return result
}

func majorityElement_(nums []int) int {
	count := 1
	majority := nums[0]
	for i := 1; i < len(nums); i++ {

		if majority == nums[i] {
			count++
		} else {
			count--
		}

		if count == 0 {
			majority = nums[i]
			count = 1
		}
	}

	return majority
}

func main() {
	testCases := []struct {
		expected int
		arr      []int
	}{
		{3, []int{3, 2, 3}},
		{2, []int{2, 2, 1, 1, 1, 2, 2}},
	}

	start := time.Now()
	for _, testCase := range testCases {
		result := majorityElement(testCase.arr)
		fmt.Printf("testPassed?= %v\n", testCase.expected == result)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken majorityElement: %s\n\n", elapsed)

	start = time.Now()
	for _, testCase := range testCases {
		result := majorityElement_(testCase.arr)
		fmt.Printf("testPassed?= %v\n", testCase.expected == result)
	}
	elapsed = time.Since(start)
	fmt.Printf("Time taken majorityElement_: %s\n\n", elapsed)
}
