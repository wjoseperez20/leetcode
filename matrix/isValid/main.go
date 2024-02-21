package main

import (
	"fmt"
	"time"
)

func isValidSudoku_(board [][]byte) bool {
	//Edge Case
	if len(board) < 9 {
		return false
	}

	cols := make([]map[byte]bool, 9)
	rows := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		cols[i] = make(map[byte]bool)
		rows[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			// Check row
			if _, exists := rows[r][board[r][c]]; exists {
				return false
			}

			// Check Column
			if _, exists := cols[c][board[r][c]]; exists {
				return false
			}

			// Check 3x3
			if _, exists := squares[(r/3)*3+c/3][board[r][c]]; exists {
				return false
			}

			rows[r][board[r][c]] = true
			cols[c][board[r][c]] = true
			squares[(r/3)*3+c/3][board[r][c]] = true
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	border := len(board)

	//Edge Case
	if len(board) < 9 {
		return false
	}

	// Check Row and Column
	for i := 0; i < border; i++ {
		rowCheck := make(map[byte]int)
		colCheck := make(map[byte]int)

		for j := 0; j < border; j++ {

			// Check Row
			if board[i][j] != '.' {
				rowCheck[board[i][j]]++
				if rowCheck[board[i][j]] == 2 {
					return false
				}
			}

			// Check Column
			if board[j][i] != '.' {
				colCheck[board[j][i]]++
				if colCheck[board[j][i]] == 2 {
					return false
				}
			}
		}
	}

	// Check 3x3
	for startRow := 0; startRow < len(board); startRow += 3 {
		for startCol := 0; startCol < len(board); startCol += 3 {

			subgridCheck := make(map[byte]int)
			for i := startRow; i < startRow+3; i++ {
				for j := startCol; j < startCol+3; j++ {

					if board[i][j] != '.' {
						subgridCheck[board[i][j]]++
						if subgridCheck[board[i][j]] == 2 {
							return false
						}
					}
				}
			}
		}
	}

	return true
}

func main() {
	testCases := []struct {
		board    [][]byte
		expected bool
	}{
		{[][]byte{}, false},
		{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}, true},
		{[][]byte{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}, false},
	}

	start := time.Now()
	for _, tc := range testCases {
		actual := isValidSudoku(tc.board)
		if actual != tc.expected {
			fmt.Printf("Fail: expected=%v, actual=%v\n", tc.expected, actual)
		}
	}
	fmt.Println("Elapsed Time: ", time.Since(start))

	start = time.Now()
	for _, tc := range testCases {
		actual := isValidSudoku_(tc.board)
		if actual != tc.expected {
			fmt.Printf("Fail: expected=%v, actual=%v\n", tc.expected, actual)
		}
	}
	fmt.Println("Elapsed Time: ", time.Since(start))
}
