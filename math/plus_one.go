package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}

// O(n)
// https://leetcode.com/problems/plus-one
func plusOne(digits []int) []int {
	current := len(digits) - 1
	digits[current]++
	for digits[current] == 10 {
		digits[current] = 0
		if current == 0 {
			digits = append([]int{1}, digits...)
			break
		} else {
			current--
			digits[current]++
		}
	}
	return digits
}
