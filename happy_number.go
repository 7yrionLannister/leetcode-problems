package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isHappy(19))
}

// https://leetcode.com/problems/happy-number
// O(log n)
func isHappy(n int) bool {
	results := make(map[int]int)
	for {
		_, present := results[n]
		if present {
			// we reached a point where we are going to loop over what we already found to realize this is not a happy number in the end
			return false
		}
		result := getDigitsSumOfSquares(n, results)
		if result == 1 {
			return true
		}
		n = result
	}
}

func getDigitsSumOfSquares(n int, results map[int]int) int {
	stringDigits := strings.Split(strconv.Itoa(n), "")
	sumOfSquares := 0
	for _, digit := range stringDigits {
		num, _ := strconv.Atoi(digit)
		result, present := results[num]
		if !present {
			result = num * num
			results[num] = result
		}
		sumOfSquares += result
	}
	return sumOfSquares
}
