package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0

	for left < right {
		// Calculate the area
		area := min(height[left], height[right]) * (right - left)

		// Update the result
		if area > res {
			res = area
		}

		// Move the pointer that has the smaller height
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, tc := range testCases {
		actual := maxArea(tc.input)
		if actual != tc.expected {
			fmt.Printf("FAIL: maxArea(%v): expected %v, actual %v\n", tc.input, tc.expected, actual)
		}
	}
}
