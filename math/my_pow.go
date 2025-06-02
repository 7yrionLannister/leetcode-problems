package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(-1, -2147483648))
}

// se podria optimizar cacheando resultados
// e.g. 2^16=((((2*2)*2*2)*2*2*2*2)*2*2*2*2*2*2*2*2): pasa de 16 multiplicaciones a tan solo 4
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
