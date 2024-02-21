package main

import (
	"fmt"
	"time"
)

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Printf("%v ---- bubbleSort %v\n", arr, arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}

	fmt.Printf("%v ---- insertionSort %v\n", arr, arr)
}

func insertionSortReverse(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] < arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}

	fmt.Printf("%v ---- insertionSort %v\n", arr, arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}

	fmt.Printf("%v ---- selectionSort %v\n", arr, arr)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2

	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			k++
			i++
		} else {
			result[k] = right[j]
			k++
			j++
		}
	}

	for i < len(left) {
		result[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		result[k] = right[j]
		k++
		j++
	}

	return result
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left, right []int
	for _, val := range arr[1:] {
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	testCases := [][]int{
		{6, 7, 3, 2, 1, 10},
		{5, 6, 4},
		{5, 6, 4, 3, 2, 1},
		{5, 6, 4, 3, 2, 1, 0},
		{5, 6, 4, 3, 2, 1, 0, -1},
		{5, 6, 4, 3, 2, 1, 0, -1, -2},
		{5, 6, 4, 3, 2, 1, 0, -1, -2, -3},
		{'e', 'd', 'c', 'b', 'a'},
		{'e', 'd', 'c', 'b', 'a', 'z', 'y', 'x', 'w', 'v'},
		{'e', 'd', 'c', 'b', 'a', 'z', 'y', 'x', 'w', 'v', 'u', 't'},
	}

	for _, val := range testCases {
		valCopy := make([]int, len(val))
		copy(valCopy, val)
		start := time.Now()
		bubbleSort(valCopy)
		elapsed := time.Since(start)
		fmt.Printf("bubbleSort took %s\n", elapsed)

		copy(valCopy, val)
		start = time.Now()
		selectionSort(valCopy)
		elapsed = time.Since(start)
		fmt.Printf("selectionSort took %s\n", elapsed)

		copy(valCopy, val)
		start = time.Now()
		insertionSort(valCopy)
		elapsed = time.Since(start)
		fmt.Printf("insertionSort took %s\n", elapsed)

		copy(valCopy, val)
		start = time.Now()
		insertionSortReverse(valCopy)
		elapsed = time.Since(start)
		fmt.Printf("insertionSortReverse took %s\n", elapsed)

		copy(valCopy, val)
		start = time.Now()
		quickSort(valCopy)
		elapsed = time.Since(start)
		fmt.Printf("quickSort took %s\n", elapsed)

		copy(valCopy, val)
		start = time.Now()
		mergeSort(valCopy)
		elapsed = time.Since(start)
		fmt.Printf("mergeSort took %s\n", elapsed)
	}
}
