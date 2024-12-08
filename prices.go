package main

import "fmt"

var maxProfitNum = 0

func main() {
	array := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(array))
	array = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(array))
	array = []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	fmt.Println(maxProfit(array))
	array = []int{2, 4, 1, 7, 11}
	fmt.Println(maxProfit(array))
	array = []int{2, 4, 1}
	fmt.Println(maxProfit(array))
	array = []int{8, 6, 4, 3, 3, 2, 3, 5, 8, 3, 8, 2, 6}
	fmt.Println(maxProfit(array))
}

func maxProfit(prices []int) int {
	maxProfitNum = 0
	maxProfitRecursive(prices, 0, len(prices)-1)
	return maxProfitNum
}

func maxProfitRecursive(prices []int, i, j int) (int, int, int) {
	var minNum, maxNum, diff int
	if j == i {
		pI := prices[i]
		minNum, maxNum, diff = pI, pI, 0
	} else if j == i+1 {
		minNum, maxNum, diff = baseCase(prices, i, j)
	} else {
		mid := (j + i) / 2
		minLeft, maxLeft, diffLeft := maxProfitRecursive(prices, i, mid)
		minRight, maxRight, diffRight := maxProfitRecursive(prices, mid+1, j)

		diff = maxRight - minLeft
		minNum = min(minLeft, minRight)
		maxNum = max(maxLeft, maxRight)
		diff = max(diff, diffLeft, diffRight)
	}
	maxProfitNum = max(maxProfitNum, diff)
	return minNum, maxNum, diff
}

func baseCase(prices []int, i, j int) (int, int, int) {
	priceI := prices[i]
	priceJ := prices[j]
	var minNum, maxNum int
	if priceI < priceJ {
		minNum = priceI
		maxNum = priceJ
	} else {
		minNum = priceJ
		maxNum = priceI
	}
	return minNum, maxNum, priceJ - priceI
}
