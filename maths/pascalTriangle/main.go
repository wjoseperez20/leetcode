package main

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func main() {
	testCases := []struct {
		numRows  int
		expected [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
		{2, [][]int{{1}, {1, 1}}},
	}

	for _, c := range testCases {
		got := generate(c.numRows)
		for i := 0; i < len(got); i++ {
			for j := 0; j < len(got[i]); j++ {
				if got[i][j] != c.expected[i][j] {
					println("Failed")
					return
				}
			}
		}

		println("Passed")
	}
}
