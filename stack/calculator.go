package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("(7)-(0)+(4)"))
	fmt.Println(calculate("1-(     -2)"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate("123"))
	fmt.Println(calculate("1+(2+3)+4"))
	fmt.Println(calculate("1+2+3+4"))
}

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)

	return calculateRecursive(s)
}

func calculateRecursive(s string) int {
	runningCalculation := 0
	n := len(s)
	prevOperator := '+'
	for i := 0; i < n; i++ {
		r := rune(s[i])
		if r == '(' {
			finish := getClosingParenthesesIndex(s, i)
			number := calculateRecursive(s[i+1 : finish])
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

func getClosingParenthesesIndex(s string, start int) int {
	stack := new(Stack[rune])
	n := len(s)
	for start < n {
		r := rune(s[start])
		if r == ')' {
			stack.Pop()
			if stack.IsEmpty() {
				return start
			}
		} else if r == '(' {
			stack.Push(r)
		}
		start++
	}
	return start
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
