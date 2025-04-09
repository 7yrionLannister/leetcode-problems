package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/"))
}

// O(n)
// https://leetcode.com/problems/simplify-path
func simplifyPath(path string) string {
	stack := new(Stack[string])
	splitPath := strings.Split(path, "/")
	for _, element := range splitPath {
		if element == "" || element == "." {
			continue
		}
		if element == ".." {
			stack.Pop()
		} else {
			stack.Push(element)
		}
	}
	if stack.IsEmpty() {
		return "/"
	}
	last, _ := stack.Pop()
	resultPath := *last
	for !stack.IsEmpty() {
		last, _ = stack.Pop()
		resultPath = *last + "/" + resultPath
	}
	return "/" + resultPath
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
