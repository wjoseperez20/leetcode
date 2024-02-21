package main

import "fmt"

// Best
func romanToInt_(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue := romanValues[s[i]]
		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}
		prevValue = currentValue
	}

	return result
}

// Solution Wilmer Perez
func romanToInt(s string) int {
	// given
	var result int
	var boolean bool
	dictionary := make(map[string]int)
	dictionary["I"] = 1
	dictionary["V"] = 5
	dictionary["X"] = 10
	dictionary["L"] = 50
	dictionary["C"] = 100
	dictionary["D"] = 500
	dictionary["M"] = 1000
	dictionary["IV"] = 4
	dictionary["IX"] = 9
	dictionary["XL"] = 40
	dictionary["XC"] = 90
	dictionary["CD"] = 400
	dictionary["CM"] = 900

	// Then
	for i := len(s) - 1; i >= 0; i-- {
		if i-1 >= 0 {
			_, boolean = dictionary[string(s[i-1])+string(s[i])]

			if boolean {
				result += dictionary[string(s[i-1])+string(s[i])]
				i--
			}
		} else {
			boolean = false
		}

		if !boolean {
			result += dictionary[string(s[i])]
		}
	}

	return result
}

func main() {
	testCases := []string{"III", "MCMXCIV", "LVIII", "MCMXCIV"}

	for _, value := range testCases {
		fmt.Printf("El valor de %v es %d\n", value, romanToInt_(value))
	}
}
