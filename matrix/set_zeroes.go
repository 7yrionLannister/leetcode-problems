package main

func setZeroes(matrix [][]int) {
	// rows := len(matrix)
	// cols := len(matrix[0])
	rowsMap := make(map[int]bool)
	colsMap := make(map[int]bool)
	for iR, row := range matrix {
		for iC, item := range row {
			if item == 0 {
				rowsMap[iR] = true
				colsMap[iC] = true
			}
		}
	}
	for row := range rowsMap {
		for i := range len(matrix[0]) {
			matrix[row][i] = 0
		}
	}
	for col := range colsMap {
		for i := range len(matrix) {
			matrix[i][col] = 0
		}
	}
}
