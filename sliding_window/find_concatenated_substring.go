package main

import "fmt"

// https://leetcode.com/problems/substring-with-concatenation-of-all-words
func main() {
	fmt.Println(findSubstring("bcabbcaabbccacacbabccacaababcbb", []string{"c", "b", "a", "c", "a", "a", "a", "b", "c"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func findSubstring(s string, words []string) []int {
	n := len(s)
	wordsMap := make(map[string]int)
	numberOfWords := len(words)
	wordLength := len(words[0])
	windowLength := wordLength * numberOfWords
	for _, word := range words {
		wordsMap[word]++
	}

	uniqueWords := len(wordsMap)
	i := 0
	indices := make([]int, 0, len(s)/windowLength)
	for j := windowLength - 1; j < n; j++ {
		runningWordsMap := make(map[string]int, uniqueWords)
		match := true
		for k := i; k <= j && match; k += wordLength {
			wordEnd := k + wordLength
			word := s[k:wordEnd]
			targetWordCount := wordsMap[word]
			currentCount := runningWordsMap[word]
			if currentCount >= targetWordCount {
				match = false
			} else {
				runningWordsMap[word]++
			}
		}
		if match {
			indices = append(indices, i)
		}
		i++
	}
	return indices
}
