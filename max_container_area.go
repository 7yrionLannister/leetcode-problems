package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 2, 4, 3}))
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}

func maxArea(height []int) int {
	n := len(height)
	indexLeft := 0
	indexRight := n - 1
	maxArea := 0
	for indexLeft < indexRight {
		maxArea = max(maxArea, area(height, indexLeft, indexRight))
		if height[indexLeft] < height[indexRight] {
			indexLeft++
		} else {
			indexRight--
		}
	}
	return maxArea
}

func area(height []int, indexLeft, indexRight int) int {
	minHeight := min(height[indexLeft], height[indexRight])
	area := minHeight * (indexRight - indexLeft)
	return area
}
