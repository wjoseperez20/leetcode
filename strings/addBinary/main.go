package main

import "fmt"

func addBinary(a string, b string) string {
	var result []rune
	flag := false

	if len(a) == len(b) {
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] == b[i] {
				result = append([]rune{'0'}, result...)

				if i == 0 {
					result = append([]rune{'1'}, result...)
					break
				}

				if a[i] == '1' && !flag {
					flag = true
				}

			} else {
				if flag {
					result = append([]rune{'0'}, result...)
				} else {
					result = append([]rune{'1'}, result...)
					flag = false
				}

			}
		}
	}

	return string(result)
}

func addBinary_(a string, b string) string {
	i, j, carry, res := len(a)-1, len(b)-1, 0, ""
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		carry = sum / 2
		res = fmt.Sprintf("%d", sum%2) + res
	}
	if carry != 0 {
		res = "1" + res
	}
	return res
}

func main() {
	testCases := []struct {
		a        string
		b        string
		expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, testCase := range testCases {
		result := addBinary_(testCase.a, testCase.b)
		fmt.Printf("result: %s, TestPassed?= %v\n", result, result == testCase.expected)
	}
}
