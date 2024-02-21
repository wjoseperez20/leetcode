package main

import (
	"fmt"
	"time"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Create a map to store the numbers
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}

	// Iterate through the map
	var max, count int
	for k, _ := range m {
		// If the number is not the start of a sequence, continue
		if m[k-1] {
			continue
		}

		// Reset the count
		count = 0

		// Iterate through the map and increment the count
		for m[k] {
			count++
			k++
		}

		// Update the max
		if count > max {
			max = count
		}
	}

	return max
}

func longestConsecutive_(nums []int) int {
	var result int

	if len(nums) == 0 {
		return result
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > result {
				result = currentStreak
			}
		}
	}

	return result
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 2, 0, 1}, 3},
		{[]int{1, 2, 0, 1, 3, 4, 5, 6, 7, 8, 9, 10}, 11},
	}

	start := time.Now()
	for _, tc := range testCases {
		got := longestConsecutive(tc.nums)
		if got != tc.want {
			fmt.Printf("longestConsecutive(%v) = %d; want %d\n", tc.nums, got, tc.want)
		}
	}
	fmt.Printf("longestConsecutive took %s\n", time.Since(start))

	start = time.Now()
	for _, tc := range testCases {
		got := longestConsecutive_(tc.nums)
		if got != tc.want {
			fmt.Printf("longestConsecutive_(%v) = %d; want %d\n", tc.nums, got, tc.want)
		}
	}
	fmt.Printf("longestConsecutive_ took %s\n", time.Since(start))
}
