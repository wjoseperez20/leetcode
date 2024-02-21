package main

import (
	"fmt"
	"time"
)

// Recursive
// Time complexity: O(2^n)
// Space complexity: O(n)
// Top down approach
func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

// Memoization
// Time complexity: O(n)
// Space complexity: O(n)
// Top down approach
func fibonacciMemo(n int) int {
	memo := make(map[int]int)
	return fibonacciMemoHelper(n, memo)
}

func fibonacciMemoHelper(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, ok := memo[n]; ok {
		return val
	} else {
		memo[n] = fibonacciMemoHelper(n-1, memo) + fibonacciMemoHelper(n-2, memo)
		return memo[n]
	}

}

// Tabulation
// Time complexity: O(n)
// Space complexity: O(n)
// Bottom up approach
func fibonacciTab(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]

}

func main() {
	testCases := map[int]int{
		5:  5,
		6:  8,
		7:  13,
		15: 610,
	}

	start := time.Now()
	for key, val := range testCases {
		result := fibonacci(key)
		if result != val {
			fmt.Printf("Error: Expected %d, got %d\n", val, result)
		}
	}
	fmt.Printf("Time taken fibonacci(Recursive): %v\n", time.Since(start))

	start = time.Now()
	for key, val := range testCases {
		result := fibonacciMemo(key)
		if result != val {
			fmt.Printf("Error: Expected %d, got %d\n", val, result)
		}
	}
	fmt.Printf("Time taken fibonacci(Memoization): %v\n", time.Since(start))

	start = time.Now()
	for key, val := range testCases {
		result := fibonacciTab(key)
		if result != val {
			fmt.Printf("Error: Expected %d, got %d\n", val, result)
		}
	}
	fmt.Printf("Time taken fibonacci(Tabulation): %v\n", time.Since(start))
}
