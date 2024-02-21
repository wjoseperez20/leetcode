package main

func count(nums []int) int {
	// Empty
	if len(nums) == 0 {
		return 0
	}

	// Base Case
	if len(nums) == 1 {
		return 1
	}

	// Recursion Case
	return 1 + count(nums[1:])
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: 5},
		{input: []int{1, 2, 3, 4, 5, 6}, expected: 6},
		{input: []int{1, 2, 3, 4, 5, 6, 7}, expected: 7},
		{input: []int{0}, expected: 1},
		{input: []int{}, expected: 0},
	}

	for _, t := range testCases {
		actual := count(t.input)
		if actual != t.expected {
			panic("Test failed")
		}
	}
}
