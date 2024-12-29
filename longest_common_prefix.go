package main

import "unicode/utf8"

// https://leetcode.com/problems/longest-common-prefix
func main() {
	strs := []string{"flower", "flow", "flight"}
	println(longestCommonPrefix(strs))
}

// time complexity: O(n * m), where n is the number of strings and m is the length of the shortest string
func longestCommonPrefix(strs []string) string {
	prefix := ""
	i := 0
	for advanceIndex, char := traverseCharsInStrings(strs, 0); advanceIndex; advanceIndex, char = traverseCharsInStrings(strs, i) {
		prefix += string(char)
		i++
	}
	return prefix

}

func traverseCharsInStrings(strs []string, index int) (bool, rune) {
	n := len(strs)
	advanceIndex := true
	char := '-'
	for i := 0; i < n && advanceIndex; i++ {
		str := strs[i]
		advanceIndex = index < len(str)
		if advanceIndex && char == '-' {
			char, _ = utf8.DecodeRuneInString(str[index:])
		}
		advanceIndex = advanceIndex && rune(str[index]) == char
	}
	return advanceIndex, char
}
