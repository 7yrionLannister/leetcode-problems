package main

import (
	"fmt"
	"slices"
	"strings"
)

// https://leetcode.com/problems/group-anagrams
func main() {
	anagrams := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(anagrams)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		key := getKey(str)
		anagrams[key] = append(anagrams[key], str)
	}
	result := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		result = append(result, v)
	}
	return result
}

func getKey(key string) string {
	keySlice := strings.Split(key, "")
	slices.Sort(keySlice)
	return strings.Join(keySlice, "")
}
