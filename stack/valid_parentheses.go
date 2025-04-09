package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
}

// https://leetcode.com/problems/valid-parentheses
func isValid(s string) bool {
	stack := new(Stack[rune])
	closeToOpen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, r := range s {
		openChar, present := closeToOpen[r]
		if present {
			top, err := stack.Pop()
			if err != nil || *top != openChar {
				return false
			}
		} else {
			stack.Push(r)
		}
	}
	return stack.IsEmpty()
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
