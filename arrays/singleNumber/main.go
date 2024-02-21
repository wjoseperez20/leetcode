package main

import (
	"fmt"
	"time"
)

func singleNumber_(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}

func singleNumber(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	var result int
	check := make(map[int]int)
	for _, value := range nums {
		check[value]++
	}

	for key, value := range check {
		if value == 1 {
			result = key
		}
	}

	return result
}

func main() {
	testCases := map[int][]int{
		1: {1},
		2: {1, 2, 1},
		4: {4, 1, 2, 1, 2},
	}

	start := time.Now()
	for key, value := range testCases {
		fmt.Printf("testPassed?=%v\n", singleNumber(value) == key)
	}
	fmt.Printf("Time: %v\n", time.Since(start))

	start = time.Now()
	for key, value := range testCases {
		fmt.Printf("testPassed?=%v\n", singleNumber_(value) == key)
	}
	fmt.Printf("Time: %v\n", time.Since(start))
}
