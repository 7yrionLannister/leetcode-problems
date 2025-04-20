package main

import "fmt"

func main() {
	fmt.Println(
		isValidSudoku([][]byte{
			{5, 3, 0, 0, 7, 0, 0, 0, 0},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 6, 0},
			{8, 0, 0, 0, 6, 0, 0, 0, 3},
			{4, 0, 0, 8, 0, 3, 0, 0, 1},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 6, 0, 0, 0, 0, 2, 8, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 5},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		}),
	)
	fmt.Println(
		isValidSudoku([][]byte{
			{8, 3, 0, 0, 7, 0, 0, 0, 0},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 6, 0},
			{8, 0, 0, 0, 6, 0, 0, 0, 3},
			{4, 0, 0, 8, 0, 3, 0, 0, 1},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 6, 0, 0, 0, 0, 2, 8, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 5},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		}),
	)
}

// https://leetcode.com/problems/valid-sudoku
// O(n^2)
func isValidSudoku(board [][]byte) bool {
	rows := [9]map[byte]bool{}
	cols := [9]map[byte]bool{}
	quadrants := [3][3]map[byte]bool{}
	x := 0
	for i := range 3 {
		for j := range 3 {
			quadrants[i][j] = make(map[byte]bool, 9)

			rows[x] = make(map[byte]bool, 9)
			cols[x] = make(map[byte]bool, 9)
			x++
		}
	}
	for i, row := range board {
		qRow := i / 3
		for j, val := range row {
			_, presentRow := rows[i][val]
			_, presentCol := cols[j][val]
			qCol := j / 3
			_, presentQuadrant := quadrants[qRow][qCol][val]
			if (val != 0 && val != '.') && (presentRow || presentCol || presentQuadrant) {
				return false
			}

			rows[i][val] = true
			cols[j][val] = true
			quadrants[qRow][qCol][val] = true
		}
	}
	return true
}
