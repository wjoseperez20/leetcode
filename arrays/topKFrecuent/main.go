package main

func topKFrequent(nums []int, k int) []int {

	// Edge Case
	if len(nums) <= 1 || k == 0 {
		return nums
	}

	// Mapping
	count := make(map[int]int)
	for _, value := range nums {
		count[value]++
	}

	// Bucket
	bucket := make([][]int, len(nums)+1)
	for key, value := range count {
		bucket[value] = append(bucket[value], key)
	}

	// Solve
	var result []int
	for i := len(bucket) - 1; i > 0 && len(result) < k; i-- {
		for _, n := range bucket[i] {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}

	for _, tc := range testCases {
		got := topKFrequent(tc.nums, tc.k)
		if len(got) != len(tc.want) {
			panic("got != want")
		}
		for i := range got {
			if got[i] != tc.want[i] {
				panic("got != want")
			}
		}
	}
}
