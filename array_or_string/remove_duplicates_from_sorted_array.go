package main

import "fmt"

func main() {
	array := []int{1, 1, 1, 2, 2, 2, 3}
	removeDuplicates(array)
	fmt.Println(array)
}

// Space: O(1)
// Time: O(n)
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	writePointer := 1
	readPointer := 1
	number := nums[0]
	numberCount := 1
	for readPointer < n {
		currentRead := nums[readPointer]
		if currentRead == number {
			numberCount++
		} else {
			number = currentRead
			numberCount = 1
		}
		nums[writePointer] = currentRead
		if numberCount > 2 {
			if currentRead != number {
				writePointer++
			}
		} else {
			writePointer++
		}
		readPointer++
	}
	return writePointer
}
