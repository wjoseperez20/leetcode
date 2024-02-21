package main

func sum(nums []int) int {
	// Empty
	if len(nums) == 0 {
		return 0
	}

	// Base Case
	if len(nums) == 1 {
		return nums[0]
	}

	// Recursion Case
	return nums[0] + sum(nums[1:])
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: 15},
		{input: []int{1, 2, 3, 4, 5, 6}, expected: 21},
		{input: []int{1, 2, 3, 4, 5, 6, 7}, expected: 28},
		{input: []int{0}, expected: 0},
		{input: []int{}, expected: 0},
	}

	for _, t := range testCases {
		actual := sum(t.input)
		if actual != t.expected {
			panic("Test failed")
		}
	}
}
