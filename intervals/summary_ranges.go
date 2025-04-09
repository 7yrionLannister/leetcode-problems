package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{-2147483648, -2147483647, 2147483647}))
	fmt.Println(summaryRanges([]int{}))
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{-1}))
	fmt.Println(summaryRanges([]int{0, 1}))
}

// https://leetcode.com/problems/summary-ranges
func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return []string{}
	}
	left := 0
	right := 0
	intervals := make([]string, 0, n)
	for i := 1; i < n; i++ {
		if nums[i] != nums[right]+1 {
			intervals = append(intervals, getRange(nums[left], nums[right]))
			left = i
		}
		right = i
	}
	if right == 0 && len(intervals) == 0 {
		intervals = append(intervals, getRange(nums[right], nums[right]))
	} else if right == n-1 {
		intervals = append(intervals, getRange(nums[left], nums[right]))
	}
	return intervals
}

func getRange(i, j int) string {
	if i == j {
		return strconv.Itoa(i)
	}
	return fmt.Sprintf("%v->%v", i, j)
}
