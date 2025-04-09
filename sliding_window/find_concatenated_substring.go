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

	i := 0
	indices := []int{}
	for j := windowLength - 1; j < n; j++ {
		runningWordsMap := make(map[string]int)
		match := true
		for k := i; k <= j && match; k += wordLength {
			wordEnd := k + wordLength
			word := s[k:wordEnd]
			targetWordCound := wordsMap[word]
			currentCount := runningWordsMap[word]
			if currentCount > targetWordCound {
				match = false
			} else {
				runningWordsMap[word]++
				if currentCount+1 > targetWordCound {
					match = false
				}
			}
		}
		if match {
			indices = append(indices, i)
		}
		i++
	}
	return indices
}
