package main

import "fmt"

func climbStairs_(n int) int {
	if n == 1 {
		return n
	}

	//Fill Slice
	var slice []int
	for idx := 1; idx < n; idx++ {
		slice = append(slice, idx)
	}

	// Check
	idx := 0
	aux := 1
	steps := 1
	for idx < n {

		if (slice[idx] + slice[idx]) == n {
			steps++
		}

		if (slice[idx]+slice[aux]) == n && aux != idx {
			steps++
		}

		if aux == n {
			aux = 0
			idx++
		} else {
			aux++
		}
	}

	return steps

}

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	testCases := []struct {
		input int
		want  int
	}{
		{input: 8, want: 34},
		{input: 2, want: 2},
		{input: 3, want: 3},
		{input: 4, want: 5},
	}

	for _, tc := range testCases {
		got := climbStairs(tc.input)
		if got != tc.want {
			fmt.Printf("climbStairs(%d) = %d, want : %d.\n", tc.input, got, tc.want)
		}
	}
}
