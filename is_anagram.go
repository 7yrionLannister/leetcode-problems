package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)
	for _, r := range s {
		sMap[r]++
	}
	tMap := make(map[rune]int)
	for _, r := range t {
		tMap[r]++
	}
	for k, v := range sMap {
		if v != tMap[k] {
			return false
		}
	}
	return true
}
