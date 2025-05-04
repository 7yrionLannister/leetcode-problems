package main

import (
	"fmt"
)

func main() {
	// res := addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})
	res := addTwoNumbers(&ListNode{}, &ListNode{})
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/add-two-numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ref := new(ListNode)
	zero := &ListNode{
		Val: 0,
	}
	result := ref
	var prevRef *ListNode
	for {
		if l1 == nil {
			l1 = zero
		}
		if l2 == nil {
			l2 = zero
		}
		if l1 == l2 {
			// both ended
			break
		}
		prevRef = ref
		ref = addDigit(l1, l2, ref)
		l1 = l1.Next
		l2 = l2.Next
	}
	if prevRef.Next.Val == 0 {
		prevRef.Next = nil
	}
	return result
}

func addDigit(l1, l2, result *ListNode) *ListNode {
	val := l1.Val + l2.Val + result.Val
	result.Val = val
	result.Next = new(ListNode)
	if val >= 10 {
		result.Val = val - 10
		result.Next.Val = 1
	}
	result = result.Next
	return result
}
