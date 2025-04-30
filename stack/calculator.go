package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(10, calculate("1+2+3+4"))
	fmt.Println(10, calculate("1+(2+3)+4"))
	fmt.Println(6, calculate("(1+(2+(3)))"))
	fmt.Println(15, calculate("(1-(4-(5+2))-3)+(6+8)"))
	fmt.Println(23, calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(11, calculate("(7)-(0)+(4)"))
	fmt.Println(3, calculate("1-(     -2)"))
	fmt.Println(3, calculate(" 2-1 + 2 "))
	fmt.Println(123, calculate("123"))
}

var closingParentheses map[int]int

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	closingParentheses = getClosingParenthesesIndexMap(s)
	return calculateRecursive(s, 0, len(s)-1)
}

func calculateRecursive(s string, start, end int) int {
	runningCalculation := 0
	prevOperator := '+'
	for i := start; i <= end; i++ {
		r := rune(s[i])
		if r == '(' {
			finish := closingParentheses[i]
			number := calculateRecursive(s, i+1, finish)
			runningCalculation += operand(prevOperator, number)
			i = finish
		} else if isDigit(r) {
			finish := getIndexOfLastDigit(s, i)
			number, _ := strconv.Atoi(s[i:finish])
			i = finish - 1
			runningCalculation += operand(prevOperator, number)
		} else {
			prevOperator = r
		}
	}
	return runningCalculation
}

func getClosingParenthesesIndexMap(s string) map[int]int {
	parenthesesStack := new(Stack[rune])
	indicesStack := new(Stack[int])
	indicesMap := make(map[int]int)
	n := len(s)
	for i := 0; i < n; i++ {
		r := rune(s[i])
		if r == ')' {
			parenthesesStack.Pop()
			index, _ := indicesStack.Pop()
			indicesMap[*index] = i
		} else if r == '(' {
			parenthesesStack.Push(r)
			indicesStack.Push(i)
		}
	}
	return indicesMap
}

func operand(operator rune, number int) int {
	switch operator {
	case '+':
		return number
	default:
		return -number
	}
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func getIndexOfLastDigit(s string, start int) int {
	n := len(s)
	for start < n {
		if !isDigit(rune(s[start])) {
			return start
		}
		start++
	}
	return n
}

type Stack[T any] struct {
	array []T
}

func (s *Stack[T]) Push(elements ...T) {
	s.array = append(s.array, elements...)
}

func (s *Stack[T]) Pop() (*T, error) {
	n := len(s.array)
	if n == 0 {
		return nil, errors.New("empty stack")
	}
	element := s.array[n-1]
	s.array = s.array[:n-1]
	return &element, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.array) == 0
}
