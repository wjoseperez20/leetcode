package main

import "time"

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -low - 1
}

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func binarySearchRecursive(arr []int, target int, low int, high int) int {
	if low > high {
		return -low - 1
	}
	mid := (low + high) / 2
	if target == arr[mid] {
		return mid
	} else if target < arr[mid] {
		return binarySearchRecursive(arr, target, low, mid-1)
	} else {
		return binarySearchRecursive(arr, target, mid+1, high)
	}
}

func main() {
	arr := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{5, 6, 4, 3, 2, 1},
	}
	for _, v := range arr {
		start := time.Now()
		println(binarySearch(v, 3))
		elipsed := time.Since(start)
		println("binarySearch took", elipsed)

		start = time.Now()
		println(linearSearch(v, 3))
		elipsed = time.Since(start)
		println("linearSearch took", elipsed)

		start = time.Now()
		println(binarySearchRecursive(v, 3, 0, len(v)-1))
		elipsed = time.Since(start)
		println("binarySearchRecursive took", elipsed)
	}
}
