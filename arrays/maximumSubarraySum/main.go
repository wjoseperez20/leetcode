package main

func maximumSubarraySum_(nums []int, k int) int64 {
	var result, sum int64
	check := make(map[int]bool)

	for i, j := 0, 0; j < len(nums); j++ {

		for check[nums[j]] {
			delete(check, nums[i])
			sum -= int64(nums[i])
			i++
		}

		check[nums[j]] = true
		sum += int64(nums[j])

		if j-i+1 == k {
			if sum > result {
				result = sum
			}

			delete(check, nums[i])
			sum -= int64(nums[i])
			i++
		}
	}

	return result
}

func maximumSubarraySum(nums []int, k int) int64 {
	var result int64

	for i := 0; i < len(nums)-k+1; i++ {

		var sum int64
		check := make(map[int]bool)

		//slide Window
		for j := i; j < i+k; j++ {

			if check[nums[j]] {
				sum = 0
				break
			}

			check[nums[j]] = true
			sum += int64(nums[j])

		}

		result = max(sum, result)

	}

	return result
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func main() {
	testCases := []struct {
		input    []int
		k        int
		expected int
	}{
		{[]int{1, 5, 4, 2, 9, 9, 9}, 3, 15},
	}

	for _, v := range testCases {
		actual := maximumSubarraySum(v.input, v.k)
		if actual != int64(v.expected) {
			println("Expected:", v.expected, "but got:", actual)
		}
	}
}
