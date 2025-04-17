package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("false ", isSubsequence("aaaaaa", "bbaaaa"))
	fmt.Println("true  ", isSubsequence("abc", "ahbgdc"))
	fmt.Println("false ", isSubsequence("b", "c"))
	fmt.Println("true  ", isSubsequence("", "ahbgdc"))
	fmt.Println("false ", isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	i := 0
	nT := len(t)
	nS := len(s)
	count := 0
	for _, r := range s {
		for i < nT && count < nS {
			char, _ := utf8.DecodeRuneInString(t[i:]) // t[i] if there are not strange characters, but use utf8 if not sure
			i++
			if char == r {
				count++
				break
			}
		}
		if i >= nT || count == nS {
			break // microoptimization, if done traversing t, stop traversing s, if count==nS return true below
		}
	}
	return count == nS
}
