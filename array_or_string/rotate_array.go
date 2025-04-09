package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	// var nums = []int{1, 2, 3, 4, 5, 6, 7}
	// var nums = []int{1, 2, 3, 4, 5, 6, 7}
	// var nums = []int{-1, -100, 3, 99}
	// var nums = []int{1}
	// var nums = []int{1, 2}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	n := len(nums)
	startIndex := 0
	replaceMe := nums[startIndex]
	for i, count := k%n, 0; count < n; i, count = (i+k)%n, count+1 {
		newReplaceMe := nums[i]
		nums[i] = replaceMe
		replaceMe = newReplaceMe
		if startIndex == i {
			startIndex = (startIndex + 1) % n
			replaceMe = nums[startIndex]
			i++
		}
	}
}
