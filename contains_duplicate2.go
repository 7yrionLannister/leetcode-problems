package main

// O(n)
// https://leetcode.com/problems/contains-duplicate-ii/description/?envType=study-plan-v2&envId=top-interview-150
func containsNearbyDuplicate(nums []int, k int) bool {
	numsIndices := make(map[int][]int)
	for i, v := range nums {
		numsIndices[v] = append(numsIndices[v], i)
		numsCurrent := numsIndices[v]
		n := len(numsCurrent)
		if n >= 2 && abs(i, numsCurrent[n-2]) <= k {
			return true
		}
	}
	return false
}

func abs(a, b int) int {
	n := a - b
	if n < 0 {
		n *= -1
	}
	return n
}
