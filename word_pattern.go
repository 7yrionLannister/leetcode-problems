package main

import "strings"

// https://leetcode.com/problems/word-pattern/
func wordPattern(pattern string, s string) bool {
	letterToWord := make(map[rune]string)
	wordToLetter := make(map[string]rune)
	sSplit := strings.Fields(s)
	if len(pattern) != len(sSplit) {
		return false
	}
	for i, r := range pattern {
		sWord := sSplit[i]
		wordForLetter, wordPresent := letterToWord[r]
		letterForWord, letterPresent := wordToLetter[sWord]
		if (wordPresent || letterPresent) && (wordForLetter != sWord || letterForWord != r) {
			return false
		}
		letterToWord[r] = sWord
		wordToLetter[sWord] = r
	}
	return true
}
