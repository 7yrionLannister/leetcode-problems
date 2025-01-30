package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("bba", "ab"))
}

func minWindow(s string, t string) string {
	sn := len(s)
	tn := len(t)
	minSubArrayLength := math.MaxInt64
	sum := 0
	iAnswer, jAnswer := 0, sn-1

	target := 0
	targetMap := make(map[rune]int)
	for _, r := range t {
		nr := normalizedAsciiValue(r)
		target += int(nr)
		targetMap[nr]++
	}

	i := 0
	for j, r := range s {
		nr := normalizedAsciiValue(r)
		_, present := targetMap[nr]
		if present {
			sum += int(nr)
		}
		for sum >= target && minSubArrayLength >= tn {
			currentSubArrayLength := j - i + 1
			if currentSubArrayLength <= minSubArrayLength {
				minSubArrayLength = currentSubArrayLength
				iAnswer, jAnswer = i, j
			}
			nr = normalizedAsciiValue(rune(s[i]))
			_, present = targetMap[nr]
			if present {
				sum -= int(nr)
			}
			i++
		}
	}
	if minSubArrayLength == math.MaxInt64 { // subarray not found
		// no subarray summed up to target or more, so minSubArrayLength == math.MaxInt64 and we have to set it to 0
		return ""
	}
	return s[iAnswer : jAnswer+1]
}

// maps lowercase characters 6 places to the left so that they are contiguous to the uppercase letters, in a sequence
func normalizedAsciiValue(r rune) rune {
	if r >= 'a' {
		r -= 6
	}
	return r
}
