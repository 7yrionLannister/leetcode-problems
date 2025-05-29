package main

import (
	"errors"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack
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

type NumberWithNode struct {
	Node   TreeNode
	Number string
}

// https://leetcode.com/problems/sum-root-to-leaf-numbers
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	traversalStack := new(Stack[NumberWithNode])
	traversalStack.Push(NumberWithNode{*root, ""})
	total := 0
	for !traversalStack.IsEmpty() {
		current, _ := traversalStack.Pop()
		current.Number += strconv.Itoa(current.Node.Val)
		if current.Node.Left == nil && current.Node.Right == nil {
			// leaf
			num, _ := strconv.Atoi(current.Number)
			total += num
			continue
		}
		if current.Node.Left != nil {
			traversalStack.Push(NumberWithNode{*current.Node.Left, current.Number})
		}
		if current.Node.Right != nil {
			traversalStack.Push(NumberWithNode{*current.Node.Right, current.Number})
		}
	}
	return total
}
