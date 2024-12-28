package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/valid-palindrome/description/?envType=study-plan-v2&envId=top-interview-150
func main() {
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	s = strings.TrimFunc(s, func(r rune) bool {
		return !isAlphanumeric(r)
	})
	palindrome := true
	n := len(s)
	i := 0
	j := n - 1
	if n > 1 {
		for i <= j && palindrome {
			r1 := toLower(rune(s[i]))
			r2 := toLower(rune(s[j]))
			if r1 != r2 {
				palindrome = false
			}
			moveUntilNextLetter(&i, 1, s)
			moveUntilNextLetter(&j, -1, s)
		}
	}
	return palindrome
}

func toLower(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func isAlphanumeric(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func moveUntilNextLetter(index *int, direction int, s string) {
	*index += direction
	for ; ; *index += direction {
		c := (rune)(s[*index])
		if isAlphanumeric(c) {
			break
		}
	}
}
