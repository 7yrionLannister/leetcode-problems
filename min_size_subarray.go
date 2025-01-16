package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/minimum-size-subarray-sum
func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

// sliding window algorithm
// O(n) as there is only one loop
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minSubArrayLength := math.MaxInt64
	i, j := 0, 0
	sum := nums[0]
	foundSum := false
	for j < n {
		if sum >= target {
			// sum is greather than or equal to target, compare (j-i+1) with minSubArrayLength and update it if it is less than it
			foundSum = true // set flag to true, to know we can trust the value of minSubArrayLength by the end of the algorithm
			currentSubArrayLength := j - i + 1
			minSubArrayLength = min(minSubArrayLength, currentSubArrayLength)
		}
		if sum > target {
			// sum exceeds target, so shrink the subarray (i++) and subtract the element left out (sum -= nums[i])
			sum -= nums[i]
			i++
			foundSum = true
		} else {
			// sum is below or equal to target, so increase the subarray size to include more positive integers that add up to the sum
			j++ // increase subarray size, which will either terminate the loop (if j == n) or end up in the first if conditional
			if j < n {
				sum += nums[j]
			}
		}
	}
	if !foundSum {
		// no subarray summed up to target or more, so minSubArrayLength == math.MaxInt64 and we have to set it to 0
		minSubArrayLength = 0
	}
	return minSubArrayLength
}
