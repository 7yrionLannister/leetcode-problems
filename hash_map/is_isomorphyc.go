package main

// https://leetcode.com/problems/isomorphic-strings
func isIsomorphic(s string, t string) bool {
	sToT := make(map[rune]rune)
	tToS := make(map[rune]rune)
	n := len(s) // == len(t)
	for i := 0; i < n; i++ {
		currentS := rune(s[i])
		currentT := rune(t[i])
		mappingSToT, okSToT := sToT[currentS]
		mappingTToS, okTToS := tToS[currentT]
		bothUnmmaped := !okSToT && !okTToS
		if bothUnmmaped { // if both currentS and currentT have no mapping it is safe to set the mapping
			tToS[currentT] = currentS
			sToT[currentS] = currentT
		} else { // if at least one of them has mapping, check whether they are the expected current characters of each string
			if mappingSToT != currentT || mappingTToS != currentS {
				return false
			}
		}
	}
	return true
}
