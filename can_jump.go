package main

import "fmt"

// https://leetcode.com/problems/jump-game
func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 5, 0, 0}))
}

var n int = 0

func canJump(nums []int) bool {
	n = len(nums)
	return canJumpRecursive(nums, 0, make([]bool, n))
}

func canJumpRecursive(nums []int, index int, visited []bool) bool {
	power := nums[index]
	visited[index] = true
	targetIndex := index + power
	if targetIndex >= n-1 {
		return true
	} else if visited[targetIndex] || (nums[index] == 0 && index < n-1) {
		return false
	} else {
		recursiveReturn := false
		for ; !recursiveReturn && targetIndex > index; targetIndex-- {
			if visited[targetIndex] {
				break
			}
			recursiveReturn = canJumpRecursive(nums, targetIndex, visited)
		}
		return recursiveReturn
	}
}
