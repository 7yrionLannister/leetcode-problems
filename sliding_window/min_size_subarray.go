package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/minimum-size-subarray-sum
func main() {
	fmt.Println(minSubArrayLenNestedLoop(11, []int{1, 2, 3, 4, 5}))
	fmt.Println(minSubArrayLenNestedLoop(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLenNestedLoop(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLenNestedLoop(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

// sliding window algorithm
// O(n) as there is only one loop
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minSubArrayLength := math.MaxInt64
	i, j := 0, 0
	sum := nums[0]
	for j < n {
		if sum >= target {
			// sum is greather than or equal to target, compare (j-i+1) with minSubArrayLength and update it if it is less than it
			currentSubArrayLength := j - i + 1
			minSubArrayLength = min(minSubArrayLength, currentSubArrayLength)
			if currentSubArrayLength == 1 {
				// base case, you can break out of the loop early because you know you are not going to find a better subarray
				break
			}
		}
		if sum > target {
			// sum exceeds target, so shrink the subarray (i++) and subtract the element left out (sum -= nums[i])
			sum -= nums[i]
			i++
		} else {
			// sum is below or equal to target, so increase the subarray size to include more positive integers that add up to the sum
			j++ // increase subarray size, which will either terminate the loop (if j == n) or end up in the first if conditional
			if j < n {
				sum += nums[j]
			}
		}
	}
	if minSubArrayLength == math.MaxInt64 { // subarray not found
		// no subarray summed up to target or more, so minSubArrayLength == math.MaxInt64 and we have to set it to 0
		minSubArrayLength = 0
	}
	return minSubArrayLength
}

// sliding window algorithm
// O(n), although there is a nested loop, i and j only traverse the elements once each, so the algorithm is still O(2n)==O(n)
func minSubArrayLenNestedLoop(target int, nums []int) int {
	n := len(nums)
	minSubArrayLength := math.MaxInt64
	i := 0
	sum := 0

	for j := 0; j < n && minSubArrayLength > 1; j++ /*grow window*/ {
		sum += nums[j] // add element to the right
		for sum >= target && minSubArrayLength > 1 {
			currentSubArrayLength := j - i + 1
			minSubArrayLength = min(minSubArrayLength, currentSubArrayLength)
			sum -= nums[i] // remove element from the left
			i++            // shrink window
		}
	}
	if minSubArrayLength == math.MaxInt64 { // subarray not found
		// no subarray summed up to target or more, so minSubArrayLength == math.MaxInt64 and we have to set it to 0
		minSubArrayLength = 0
	}
	return minSubArrayLength
}
