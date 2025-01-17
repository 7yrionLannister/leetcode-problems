package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters
// sliding window algorithm
// O(n), although there is a nested loop, i and j only traverse the elements once each, so the algorithm is still O(2n)==O(n)
func lengthOfLongestSubstring(s string) int {
	runeSet := make(map[rune]bool)
	maxSubStringLen := math.MinInt64
	i := 0
	for j, r := range s {
		_, present := runeSet[r]
		if !present {
			maxSubStringLen = max(maxSubStringLen, j-i+1)
		} else {
			for present {
				maxSubStringLen = max(maxSubStringLen, j-i)
				delete(runeSet, rune(s[i]))
				i++
				_, present = runeSet[r]
			}
		}
		runeSet[r] = true
	}
	if maxSubStringLen == math.MinInt64 {
		maxSubStringLen = len(s)
	}
	return maxSubStringLen
}
