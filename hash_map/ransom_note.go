package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("aa", "aba"))
}

// O(M+N)
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) { // optimization
		return false
	}
	lettersNeeded := make(map[rune]uint)
	for _, r := range ransomNote {
		lettersNeeded[r]++
	}
	for _, r := range magazine {
		count, present := lettersNeeded[r]
		if present {
			count--
			lettersNeeded[r] = count
			if count == 0 {
				delete(lettersNeeded, r)
				if len(lettersNeeded) == 0 {
					return true
				}
			}
		}
	}
	return false
}
