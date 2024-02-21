package main

import (
	"fmt"
	"reflect"
)

func rotate_(nums []int, k int) {

	if k >= 1 {
		n := len(nums) - 1
		for i := 0; i < k; i++ {
			val := nums[n]
			nums = append([]int{val}, nums[:n]...)
		}
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, tc := range testCases {
		rotate(tc.nums, tc.k)

		if reflect.DeepEqual(tc.nums, tc.expected) {
			fmt.Printf("Testcase passed: %v\n", tc.nums)
		} else {
			fmt.Printf("Testcase failed: %v\n", tc.nums)
		}

	}
}
