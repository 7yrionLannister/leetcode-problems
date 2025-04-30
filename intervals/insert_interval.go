package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(insert([][]int{{1, 5}}, []int{6, 8}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(i1, i2 []int) int {
		return i1[0] - i2[0]
	})
	n := len(intervals)
	mergedIntervals := make([][]int, 0, n)
	mergedIntervals = append(mergedIntervals, intervals[0])
	if n > 1 {
		if intervalsOverlap(intervals[0], intervals[1]) {
			mergedIntervals = append(mergedIntervals[0:0], mergeIntervals(intervals[0], intervals[1]))
		} else {
			mergedIntervals = append(mergedIntervals, intervals[1])
		}
	}
	for i := 2; i < n; i++ {
		m := len(mergedIntervals)
		if intervalsOverlap(intervals[i], mergedIntervals[m-1]) {
			mergedIntervals = append(mergedIntervals[:m-1], mergeIntervals(mergedIntervals[m-1], intervals[i]))
		} else {
			mergedIntervals = append(mergedIntervals, intervals[i])
		}
	}
	return mergedIntervals
}

func intervalsOverlap(i1, i2 []int) bool {
	return (i1[0] <= i2[1] && i1[0] >= i2[0]) || (i1[1] >= i2[0] && i1[1] <= i2[1]) ||
		(i2[0] <= i1[1] && i2[0] >= i1[0]) || (i2[1] >= i1[0] && i2[1] <= i1[1]) ||
		contains(i1, i2)
}

func contains(i1, i2 []int) bool {
	return (i1[0] >= i2[0] && i1[1] <= i2[1]) || (i2[0] >= i1[0] && i2[1] <= i1[1])
}

func mergeIntervals(i1, i2 []int) []int {
	min := min(i1[0], i2[0])
	max := max(i1[1], i2[1])
	return []int{min, max}
}
