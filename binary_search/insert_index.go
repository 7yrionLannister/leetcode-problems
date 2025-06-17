package main

// https://leetcode.com/problems/search-insert-position
func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	var mid int
	var numMid int
	for i <= j {
		mid = (j + i) / 2
		numMid = nums[mid]
		if numMid == target {
			return mid
		} else if numMid < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if numMid < target {
		mid++
	}
	return mid
}
