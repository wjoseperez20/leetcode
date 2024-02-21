package main

import (
	"log"
	"reflect"
)

func averageSlidingWindow(arr []int, k int) []float32 {
	var result []float32

	for i := 0; i < len(arr)-k+1; i++ {

		sum := 0
		for j := i; j < i+k; j++ {
			sum += arr[j]
		}

		result = append(result, float32(sum)/float32(k))
	}

	return result
}

func main() {
	testCases := []struct {
		input    []int
		k        int
		expected []float32
	}{
		{[]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5, []float32{2.2, 2.8, 2.4, 3.6, 2.8}},
	}

	for _, tc := range testCases {
		actual := averageSlidingWindow(tc.input, tc.k)
		if !reflect.DeepEqual(actual, tc.expected) {
			log.Fatalf("expected %v, but got %v", tc.expected, actual)
		}
	}
}
