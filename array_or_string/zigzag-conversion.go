package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("AB", 1))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("ABCDEFGHIJKLMN", 3))
	fmt.Println(convert("ABCDEFGHIJKLMN", 4))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	rowIndex := 0
	dir := 1
	for _, character := range s {
		rows[rowIndex] += string(character)
		rowIndex = (rowIndex + dir) % numRows
		if rowIndex == 0 && dir == -1 {
			dir = 1
		} else if rowIndex == numRows-1 && dir == 1 {
			dir = -1
		}
	}
	result := ""
	for _, row := range rows {
		result += row
	}
	return result
}
