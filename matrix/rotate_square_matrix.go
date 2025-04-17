package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 4}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
	matrix = [][]int{{1, 2}, {3, 4}}
	rotate(matrix)
	fmt.Println(matrix)
}

// https://leetcode.com/problems/rotate-image
// in-place
// O(n^2)
func rotate(matrix [][]int) {
	n := len(matrix)
	rings := n / 2
	for ring := range rings {
		ringWidth := n - 2*ring
		ringEnd := ring + ringWidth - 1
		// the rotation has to be of 90 degrees, so every ring has to be rotated its width minus one
		// times in order for the corner to be in the opposite corner
		for range ringWidth - 1 {
			saveMe := matrix[ring][ring]
			for col := ring + 1; col <= ringEnd; col++ { // top side
				temp := saveMe
				saveMe = matrix[ring][col]
				matrix[ring][col] = temp
			}
			for row := ring + 1; row <= ringEnd; row++ { // right side
				temp := saveMe
				saveMe = matrix[row][ringEnd]
				matrix[row][ringEnd] = temp
			}
			for col := ringEnd - 1; col >= ring; col-- { // bottom side
				temp := saveMe
				saveMe = matrix[ringEnd][col]
				matrix[ringEnd][col] = temp
			}
			for row := ringEnd - 1; row >= ring; row-- { // left side
				temp := saveMe
				saveMe = matrix[row][ring]
				matrix[row][ring] = temp
			}
		}
	}
}
