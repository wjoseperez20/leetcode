package main

import "fmt"

func lastTwoDigits(arr []int) int {
	var lastTwo int

	product := 1
	for _, val := range arr {
		product *= val
	}

	lastTwo = product % 100

	return lastTwo
}

func main() {
	testCases := map[int][]int{
		50: {25, 10},
		40: {2, 4, 5},
		1:  {1},
		0:  {0},
	}

	for key, val := range testCases {
		result := lastTwoDigits(val)
		fmt.Printf("For arr= %v the last two digits is %v. TestPassed?%v\n", val, result, result == key)
	}
}
