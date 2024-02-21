package main

import "fmt"

// Solution ChatGPT
func twoSum_(nums []int, target int) []int {
	numIndexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numIndexMap[complement]; found {
			return []int{index, i}
		}

		numIndexMap[num] = i
	}

	return nil // If no solution is found
}

// Solution Wilmer Perez
func twoSum(nums []int, target int) []int {
	// Given
	var result []int
	var p1, p2 int
	flag := true

	// First Condition
	if nums[p1] == target {
		result = []int{0}
	}

	p2 = len(nums) - 1
	for flag {
		lookTarget := nums[p1] + nums[p2]
		if lookTarget == target {
			flag = false
			result = []int{p1, p2}
		}

		if p2 > p1+1 {
			p2--
		} else {
			p2 = len(nums) - 1
			p1++
		}
	}

	return result
}

func main() {
	testCases := []struct {
		target int
		arr    []int
	}{
		{11, []int{4, 3, 6, 8}},
		{9, []int{2, 7, 11, 15}},
		{6, []int{3, 2, 4}},
		{6, []int{3, 3}},
		{20, []int{1, 3, 6, 8, 14, 13}},
	}

	for _, value := range testCases {
		fmt.Printf("For the nums= %v with target %v, ", value.arr, value.target)
		fmt.Printf("Solution = %v \n", twoSum(value.arr, value.target))
	}
}
