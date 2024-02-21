package main

import "fmt"

func mySqrt(x int) int {
	if x <= 0 {
		return x
	}

	check := true
	var result, count, operation int

	for check {
		operation = x / 2

		if operation > 0 {
			count++
		} else {
			check = false
		}

		x = operation
	}

	result = count / 2

	return result
}

func mySqrt_(x int) int {
	if x <= 1 {
		return x
	}

	// Initial guess for the square root
	guess := x / 2

	// Iterate until the guess converges
	for guess*guess > x {
		guess = (guess + x/guess) / 2
	}

	return guess
}

func main() {
	testCases := []struct {
		input  int
		output int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
	}

	for _, testCase := range testCases {
		result := mySqrt_(testCase.input)
		fmt.Printf("Expected: %v, got: %v , TestPassed?=%v\n", testCase.output, result, result == testCase.output)
	}
}
