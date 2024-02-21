package main

/*
We are given a matrix where each cell can have one of three values:
* 0: empty cell
* 1: a fresh apple
* 2: a rotten apple
Every hour, any fresh apple that is 4-directionally adjacent to a rotten apple becomes rotten.
Return the *minimum number of hours* that must elapse until no cell has a fresh apple. If this is impossible, return `-1`.

Example 1: [[2,1,1],[1,1,0],[0,1,1]]
It takes 4 hours for all apples to become rotten. Output should be `4`.
          x, y
				 (0,0) (0,1)
2 1 1       2 2 1       2 2 2       2 2 2       2 2 2
1 1 0  ->   2 1 0   ->  2 2 0   ->  2 2 0   ->  2 2 0
0 1 1       0 1 1       0 1 1       0 2 1       0 2 2

Example 2: [[2,1,1],[0,1,1],[1,0,1]]
In this example, the left bottom corner apple is isolated and won't get rotten ever. Therefore it should return `-1`.
2 1 1       2 2 1       2 2 2       2 2 2       2 2 2
0 1 1   ->  0 1 1   ->  0 2 1   ->  0 2 2   ->  0 2 2
1 0 1       1 0 1       1 0 1       1 0 1       1 0 2

Example 3: [[0,2]]
In this example, all apples are already rotten, therefore it will return `0`.
0 2

Example 4: [[1,1,2],[1,0,1]]
In this example, It takes 3 hours for all the apples to become rotten.

1 1 2      1 2 2     2 2 2   2 2 2
1 0 2      1 0 2     1 0 2   2 0 2
*/

import "fmt"

type coordinate struct {
	x int
	y int
}

func applesRotting_(grid [][]int) int {
	position := len(grid[0])
	coordinatesEmpty := make(map[coordinate]int)
	coordinatesFresh := make(map[coordinate]int)
	coordinatesRotten := make(map[coordinate]int)

	// Iterate through the grid
	for i := 0; i < position; i++ {
		for j := 0; j < position; j++ {

			// Empty cell
			if grid[i][j] == 0 {
				coordinatesEmpty[coordinate{i, j}]++
			}

			// Fresh apple
			if grid[i][j] == 1 {
				coordinatesFresh[coordinate{i, j}]++
			}

			// Rotten apple
			if grid[i][j] == 2 {
				coordinatesRotten[coordinate{i, j}]++
			}
		}

	}

	// If there are no fresh apples, return 0
	if len(coordinatesFresh) == 0 {
		return 0
	}

	for rotten, _ := range coordinatesRotten {
		row := rotten.x
		col := rotten.y

		// Check if there are fresh apples in the same row
		for i := 0; i < position; i++ {
			coordinatesFresh[coordinate{row, i}]--
		}

		// Check if there are fresh apples in the same column
		for j := 0; j < position; j++ {
			coordinatesFresh[coordinate{j, col}]--
		}
	}

	return 0
}

func applesRotting(grid [][]int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var queue [][]int
	fresh := 0
	hours := 0

	// Initialize the queue and count the fresh apples
	for i, row := range grid {
		for j, cell := range row {
			if cell == 2 {
				queue = append(queue, []int{i, j})
			} else if cell == 1 {
				fresh++
			}
		}
	}

	// BFS to propagate the rot
	for len(queue) != 0 {
		for range queue {
			rotten := queue[0]
			queue = queue[1:]

			for _, direction := range directions {
				x, y := rotten[0]+direction[0], rotten[1]+direction[1]
				if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 1 {
					continue
				}
				grid[x][y] = 2
				fresh--
				queue = append(queue, []int{x, y})
			}
		}

		if len(queue) != 0 {
			hours++
		}
	}

	// If there are still fresh apples, return -1
	if fresh > 0 {
		return -1
	}
	return hours
}

func main() {
	testCases := []struct {
		input  [][]int
		output int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
	}

	for _, testCase := range testCases {
		result := applesRotting(testCase.input)
		if result == testCase.output {
			fmt.Printf("Expected: %v, got: %v , TestPassed?=%v\n", testCase.output, result, result == testCase.output)
		}
	}
}
