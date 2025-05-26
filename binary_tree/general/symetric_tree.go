package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Queue
type Queue[T any] struct {
	array []T
}

func (q *Queue[T]) Queue(element T) {
	q.array = append(q.array, element)
}

func (q *Queue[T]) Unqueue() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	first := q.array[0]
	q.array = q.array[1:]
	return &first, nil
}

func (q *Queue[T]) Len() int {
	return len(q.array)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := new(Queue[*TreeNode])
	queue.Queue(root)
	for !queue.IsEmpty() {
		levelSize := queue.Len()
		mid := levelSize / 2
		mirrorStack := new(Stack[*TreeNode])
		levelIsfullOfNills := true
		for i := 0; i < levelSize; i++ {
			n, _ := queue.Unqueue()
			node := *n
			if node != root {
				if i < mid {
					mirrorStack.Push(node)
				} else {
					o, _ := mirrorStack.Pop()
					opening := *o
					if opening == nil && node == nil {
					} else if (opening != nil && node == nil) || (opening == nil && node != nil) || ((*opening).Val != node.Val) {
						return false
					}
				}
			}
			if node != nil {
				queue.Queue(node.Left)
				queue.Queue(node.Right)
				if node.Left != nil || node.Right != nil {
					levelIsfullOfNills = false
				}
			}
		}
		if !mirrorStack.IsEmpty() {
			return false
		}
		if levelIsfullOfNills {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(isSymmetric(&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}))
	fmt.Println(isSymmetric(&TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}))
	fmt.Println(isSymmetric(&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}))
}
