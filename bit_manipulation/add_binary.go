package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("1", "1"))
	fmt.Println(addBinary("1", "0"))
}

// https://leetcode.com/problems/add-binary
// Time: O(n)
func addBinary(a string, b string) string {
	pointerA := len(a) - 1
	pointerB := len(b) - 1
	var result string
	var remainder byte
	for {
		if pointerA < 0 && pointerB < 0 {
			if remainder == 1 {
				result = "1" + result
			}
			break
		}
		var digitA byte
		if pointerA >= 0 {
			digitA = a[pointerA] - 48
			pointerA--
		}
		var digitB byte
		if pointerB >= 0 {
			digitB = b[pointerB] - 48
			pointerB--
		}
		sum := digitA + digitB + remainder
		if sum > 1 {
			if sum == 2 {
				result = "0" + result
			} else {
				// 3
				result = "1" + result
			}
			remainder = 1
		} else {
			remainder = 0
			result = strconv.Itoa(int(sum)) + result
		}
	}
	return result
}
