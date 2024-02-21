package main

import (
	"fmt"
	"reflect"
)

func productExceptSelf(nums []int) []int {
	var result []int
	var idx, check, except int
	n := len(nums) - 1

	product := 1
	for check <= n {
		if idx != except {
			product *= nums[idx]
			idx++
		} else {
			idx++
		}

		if idx > n {
			result = append(result, product)
			product = 1
			idx = 0
			except++
			check++
		}
	}

	return result
}

func productExceptSelf_(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	// answer[i] contains the product of all the numbers to the left of i.
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// r contains the product of all the numbers to the right of i.
	r := 1
	for i := length - 1; i >= 0; i-- {
		// For the index 'i', r would contain the product of all numbers to the right. We update r accordingly
		answer[i] = answer[i] * r
		r *= nums[i]
	}

	return answer
}

func main() {
	testCase := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, v := range testCase {
		result := productExceptSelf_(v.input)
		if reflect.DeepEqual(result, v.output) {
			fmt.Printf("Expected: %v, got: %v\n", v.output, result)
		}
	}
}
