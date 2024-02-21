package main

import "fmt"

func isPalindrome(number int) bool {
	// Given
	if number < 0 {
		return false
	}

	condition := number
	reverseNumber := 0
	for condition > 0 {
		reverseNumber = reverseNumber*10 + condition%10
		condition /= 10
	}

	if (number - reverseNumber) == 0 {
		return true
	}

	return false
}

func main() {
	// Given
	testNumber := []int{573, 121, 895, 121, -125, 950, 555, 496}
	for _, value := range testNumber {
		fmt.Printf("The number %d is %v\n", value, isPalindrome(value))
	}

}
