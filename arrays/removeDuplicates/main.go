package main

import (
	"fmt"
	"time"
)

func removeDuplicatesEasy_(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize the pointer for the modified array
	uniquePointer := 0

	// Iterate through the array starting from the second element
	for i := 1; i < len(nums); i++ {
		// Check for duplicates
		if nums[i] != nums[uniquePointer] {
			// If not a duplicate, move the uniquePointer forward and update the element
			uniquePointer++
			nums[uniquePointer] = nums[i]
		}
	}

	// Length of the modified array (unique elements)
	return uniquePointer + 1
}

func removeDuplicatesEasy(nums []int) int {
	k := 1
	if len(nums) <= 1 {
		return k
	}

	p1 := 0
	p2 := 1
	for p2 <= len(nums)-1 {
		if nums[p1] == nums[p2] {
			p2++
		} else {
			k++
			nums[p1+1] = nums[p2]
			p1++
			p2++
		}
	}

	return k
}

func removeDuplicatesMedium(nums []int) int {
	result := 1
	if len(nums) <= 1 {
		return result
	}

	check := make(map[int]int)
	ptr := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		check[nums[i]]++

		if check[nums[i]] >= 3 {
			val := nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, val)
			i--
			ptr--
		}

		if ptr == i {
			result += i
			break
		}
	}

	return result
}

func removeDuplicatesMedium_(nums []int) int {
	count := 1
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

func main() {
	testCases := [][]int{
		{0, 0, 1, 1, 1, 1, 2, 3, 3},
	}

	start := time.Now()
	for _, value := range testCases {
		valCopy := make([]int, len(value))
		copy(valCopy, value)
		result := removeDuplicatesEasy(valCopy)
		fmt.Printf("The k for the %v is %v\n", value, result)
	}
	elapsed := time.Since(start)
	fmt.Printf("removeDuplicatesEasy function took %s\n", elapsed)

	start = time.Now()
	for _, value := range testCases {
		valCopy := make([]int, len(value))
		copy(valCopy, value)
		result := removeDuplicatesEasy_(valCopy)
		fmt.Printf("The k for the %v is %v\n", value, result)
	}
	elapsed = time.Since(start)
	fmt.Printf("removeDuplicatesEasy_ function took %s\n", elapsed)

	start = time.Now()
	for _, value := range testCases {
		valCopy := make([]int, len(value))
		copy(valCopy, value)
		result := removeDuplicatesMedium(valCopy)
		fmt.Printf("The k for the %v is %v\n", value, result)
	}
	elapsed = time.Since(start)
	fmt.Printf("removeDuplicatesMedium function took %s\n", elapsed)

	start = time.Now()
	for _, value := range testCases {
		valCopy := make([]int, len(value))
		copy(valCopy, value)
		result := removeDuplicatesMedium_(valCopy)
		fmt.Printf("The k for the %v is %v\n", value, result)
	}
	elapsed = time.Since(start)
	fmt.Printf("removeDuplicatesMedium_ function took %s\n", elapsed)
}
