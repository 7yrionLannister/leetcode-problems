package main

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println("false ", isSubsequence("aaaaaa", "bbaaaa"))
	fmt.Println("true  ", isSubsequence("abc", "ahbgdc"))
	fmt.Println("false ", isSubsequence("b", "c"))
	fmt.Println("true  ", isSubsequence("", "ahbgdc"))
	fmt.Println("false ", isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	constraints.Ordered
	i := 0
	nT := len(t)
	nS := len(s)
	count := 0
	for _, r := range s {
		for ; i < nT; i++ {
			char, _ := utf8.DecodeRuneInString(t[i:])
			if char == r {
				i++
				count++
				break
			}
		}
		if i >= nT {
			break
		}
	}
	return count == nS
}
