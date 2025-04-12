package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}

// O(n)
// https://leetcode.com/problems/spiral-matrix
func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])

	zeroIndexRow := 0
	topIndexRow := rows - 1

	zeroIndexCol := 0
	topIndexCol := cols - 1

	dirY := 0
	dirX := 1
	total := rows * cols
	spiral := make([]int, 0, total)
	for count := 0; count < total; {
		if dirX == 0 {
			if dirY == 1 {
				// down
				for rowIdx := zeroIndexRow; rowIdx <= topIndexRow; rowIdx++ {
					spiral = append(spiral, matrix[rowIdx][topIndexCol])
					count++
				}
				topIndexCol--
				dirY = 0
				dirX = -1
			} else {
				// up
				for rowIdx := topIndexRow; rowIdx >= zeroIndexRow; rowIdx-- {
					spiral = append(spiral, matrix[rowIdx][zeroIndexCol])
					count++
				}
				zeroIndexCol++
				dirY = 0
				dirX = 1
			}
		} else {
			if dirX == 1 {
				// right
				for colIdx := zeroIndexCol; colIdx <= topIndexCol; colIdx++ {
					spiral = append(spiral, matrix[zeroIndexRow][colIdx])
					count++
				}
				zeroIndexRow++
				dirY = 1
				dirX = 0
			} else {
				// left
				for colIdx := topIndexCol; colIdx >= zeroIndexCol; colIdx-- {
					spiral = append(spiral, matrix[topIndexRow][colIdx])
					count++
				}
				topIndexRow--
				dirY = -1
				dirX = 0
			}
		}
	}
	return spiral
}
