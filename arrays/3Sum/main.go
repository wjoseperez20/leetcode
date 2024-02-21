package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

func threeSum__(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i, a := range nums {
		// Skip positive integers
		if a > 0 {
			break
		}

		// Skip duplicates
		if i > 0 && a == nums[i-1] {
			continue
		}

		// Two pointers
		l, r := i+1, len(nums)-1
		for l < r {
			target := a + nums[l] + nums[r]
			if target > 0 {
				r--
			} else if target < 0 {
				l++
			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				l++
				r--
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return res
}

//----------------------------------------------------------------------------------------------

// Solution Wilmer Perez
func threeSum(nums []int) [][]int {
	var result [][]int

	lookUp := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {

				if nums[i]+nums[j]+nums[k] == 0 {

					check := bubbleSort([]int{nums[i], nums[j], nums[k]})
					if _, ok := lookUp[[3]int(check)]; !ok {
						result = append(result, check)

						lookUp[[3]int(check)] = true
					}
				}
			}
		}
	}

	return result
}

func bubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// ----------------------------------------------------------------------------------------------
// Solution ChatGPT
func threeSum_(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	// Sort the array
	quickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			low, high := i+1, len(nums)-1
			sum := 0 - nums[i]

			for low < high {
				if nums[low]+nums[high] == sum {
					result = append(result, []int{nums[i], nums[low], nums[high]})
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if nums[low]+nums[high] < sum {
					low++
				} else {
					high--
				}
			}
		}
	}

	return result
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	testCases := []struct {
		target   int
		arr      []int
		expected [][]int
	}{
		{target: 0, arr: []int{-1, 0, 1, 2, -1, -4}, expected: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{0, []int{0, 1, 1}, [][]int{}},
		{0, []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	start := time.Now()
	for idx, test := range testCases {
		result := threeSum(test.arr)

		if reflect.DeepEqual(result, test.expected) {
			fmt.Printf("Test case %d - Passed\n", idx+1)
		} else {
			fmt.Printf("Test case %d - Failed\n", idx+1)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)

	start = time.Now()
	for idx, test := range testCases {
		result := threeSum_(test.arr)

		if reflect.DeepEqual(result, test.expected) {
			fmt.Printf("Test case %d - Passed\n", idx+1)
		} else {
			fmt.Printf("Test case %d - Failed\n", idx+1)
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)

	start = time.Now()
	for idx, test := range testCases {
		result := threeSum__(test.arr)

		if reflect.DeepEqual(result, test.expected) {
			fmt.Printf("Test case %d - Passed\n", idx+1)
		} else {
			fmt.Printf("Test case %d - Failed\n", idx+1)
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
