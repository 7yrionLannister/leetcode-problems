package main

import (
	"errors"
	"strconv"
)

// O(n)
// https://leetcode.com/problems/evaluate-reverse-polish-notation
func evalRPN(tokens []string) int {
	operandsStack := new(Stack[int])
	for _, token := range tokens {
		switch token {
		case "-":
			top, secondTop := getTwoTops(operandsStack)
			res := *secondTop - *top
			operandsStack.Push(res)
		case "+":
			top, secondTop := getTwoTops(operandsStack)
			res := *secondTop + *top
			operandsStack.Push(res)
		case "*":
			top, secondTop := getTwoTops(operandsStack)
			res := *secondTop * *top
			operandsStack.Push(res)
		case "/":
			top, secondTop := getTwoTops(operandsStack)
			res := *secondTop / *top
			operandsStack.Push(res)
		default:
			num, _ := strconv.Atoi(token)
			operandsStack.Push(num)
		}
	}
	result, _ := operandsStack.Pop()
	return *result
}

func getTwoTops[T any](s *Stack[T]) (top *T, secondTop *T) {
	top, _ = s.Pop()
	secondTop, _ = s.Pop()
	return
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
