package main

import "time"

func factorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	testCases := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{5, 120},
		{10, 3628800},
	}

	start := time.Now()
	for _, testCase := range testCases {
		if res := factorial(testCase.input); res != testCase.output {
			panic("Test failed")
		}
	}
	elapsed := time.Since(start)
	println("factorial took", elapsed)
}
