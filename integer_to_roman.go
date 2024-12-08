package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(intToRoman(49))
	fmt.Println(intToRoman(3749))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(900))
}

func intToRoman(num int) string {
	romans := make(map[int]string)
	romans[1000] = "M"
	romans[500] = "D"
	romans[100] = "C"
	romans[50] = "L"
	romans[10] = "X"
	romans[5] = "V"
	romans[1] = "I"

	array := []int{1000, 500, 100, 50, 10, 5, 1}
	result := ""
	n := len(array)
	for i := 0; i < n; i++ {
		value := array[i]
		division := num / value
		firstChar := rune(strconv.Itoa(num)[0])
		if division > 0 {
			if firstChar == '4' {
				if value == 100 {
					result += "CD"
					num -= 400
				} else if value == 50 {
					result += "XL"
					num -= 40
				} else {
					result += "IV"
					num -= 4
				}
			} else if firstChar == '9' {
				if value == 500 {
					result += "CM"
					num -= 900
				} else if value == 50 {
					result += "XC"
					num -= 90
				} else {
					result += "IX"
					num -= 9
				}
			} else {
				result += repeatChar(romans[value], division)
				num -= value * division
			}
		} else if i < n-1 && num > array[i+1] && (firstChar == '4' || firstChar == '9') {
			concat := romans[value]
			if firstChar == '4' {
				if value == 500 {
					concat = "C" + concat
					num -= 400
				} else if value == 50 {
					concat = "X" + concat
					num -= 40
				} else {
					concat = "I" + concat
					num -= 4
				}
			} else if firstChar == '9' {
				if value == 1000 {
					concat = "C" + concat
					num -= 900
				} else if value == 100 {
					concat = "X" + concat
					num -= 90
				} else {
					concat = "I" + concat
					num -= 9
				}
			}
			result += concat
		}
		if num == 0 {
			break
		}
	}
	return result
}

func repeatChar(char string, times int) string {
	result := ""
	for times > 0 {
		result += char
		times--
	}
	return result
}
