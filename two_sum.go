package main

import "fmt"

// O(n)
func twoSum(nums []int, target int) []int {
	numberToIndex := make(map[int]int)
	{
		fmt.Println("")
		println("asd")
	}
	for i, v := range nums {
		index, present := numberToIndex[target-v]
		if present {
			return []int{i, index}
		}
		numberToIndex[v] = i
	}
	return []int{-1, -1}
}

//O(n^2)
/*
func twoSum(nums []int, target int) []int {
    n := len(nums)
    previouslyComputed := make(map[int]bool)
    for i, v := range nums {
        for j := i+1; j < n && !previouslyComputed[i]; j++ {
            if v + nums[j] == target {
                return []int{i, j}
            }
        }
        previouslyComputed[i] = true
    }
    return []int{-1, -1}
}
*/
