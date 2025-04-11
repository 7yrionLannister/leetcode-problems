package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{7, -9, 3, -6, 3, 5, 3, 6, -2, -5, 8, 6, -4, -6, -4, -4, 5, -9, 2, 7, 0, 0}), 4)
	fmt.Println(longestConsecutive([]int{-3, 2, 8, 5, 1, 7, -8, 2, -8, -4, -1, 6, -6, 9, 6, 0, -7, 4, 5, -4, 8, 2, 0, -2, -6, 9, -4, -1}), 7)
	fmt.Println(longestConsecutive([]int{-1, -9, -5, -2, -9, 8, -8, 1, -9, -3, -3}), 3)
	fmt.Println(longestConsecutive([]int{-7, -1, 3, -9, -4, 7, -3, 2, 4, 9, 4, -9, 8, -7, 5, -1, -7}), 4)
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), 9)
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}), 4)
}

// https://leetcode.com/problems/longest-consecutive-sequence
func longestConsecutive(nums []int) int {
	rangeMap := make(map[int]int)
	reverseRangeMap := make(map[int]int)
	for _, num := range nums {
		_, presentAsTail := rangeMap[num]
		_, presentAsHead := reverseRangeMap[num]
		if presentAsTail || presentAsHead {
			continue
		}
		prev := num - 1
		next := num + 1
		prevMapping, presentPrev := rangeMap[prev]
		if !presentPrev {
			prevMapping, presentPrev = reverseRangeMap[prev]
			if presentPrev {
				delete(reverseRangeMap, prev)
				reverseRangeMap[num] = prevMapping
				rangeMap[prevMapping] = num
				prev = prevMapping
			}
		}
		nextMapping, presentNext := rangeMap[next]
		if presentPrev && prevMapping == prev {
			rangeMap[prev] = num
			reverseRangeMap[num] = prev
			delete(reverseRangeMap, prev)
			if presentNext {
				delete(reverseRangeMap, num)
			}
			num = prev
		}
		if presentNext {
			rangeMap[num] = nextMapping
			delete(rangeMap, next)
			reverseRangeMap[nextMapping] = num
			delete(reverseRangeMap, num)
		}
		if !presentPrev && !presentNext {
			rangeMap[num] = num
			reverseRangeMap[num] = num
		}
	}
	highestRangeSize := 0
	for low, high := range rangeMap {
		highestRangeSize = max(highestRangeSize, high-low+1)
	}
	return int(highestRangeSize)
}
