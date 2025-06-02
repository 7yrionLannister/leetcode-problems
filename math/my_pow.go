package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(-1, -2147483648))
}

// O (n)
// https://leetcode.com/problems/powx-n/submissions
func myPow(x float64, n int) float64 {
	result := x
	if x != 1 && x != -1 {
		for range int((math.Abs(float64(n)) - 1)) {
			result *= x
			if result == math.Inf(1) {
				break
			}
		}
	}
	if n == 0 {
		result = 1
	} else if x == -1 && n%2 == 0 {
		result = 1
	} else if n < 0 {
		result = 1 / result
	}
	return result
}
