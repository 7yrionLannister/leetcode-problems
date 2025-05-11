package main

import "fmt"

func main() {
	// res := partition(&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}, 3)
	// res := partition(&ListNode{2, &ListNode{1, nil}}, 2)
	res := partition(&ListNode{1, nil}, 0)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n)
// Space: O(1)
// https://leetcode.com/problems/partition-list
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	init := head
	var m1s *ListNode
	var m2s *ListNode
	var m1f *ListNode
	var m2f *ListNode
	if head.Val < x {
		m1s = head
		m1f = head
		for head != nil && head.Val < x {
			m1f = head
			head = head.Next
		}
		m2s = head
		m2f = head
	} else {
		m2s = head
		m2f = head
		for head != nil && head.Val >= x {
			m2f = head
			head = head.Next
		}
		m1s = head
		m1f = head
	}

	if head != nil {
		head = head.Next
	}
	for head != nil {
		if head.Val < x {
			m1f.Next = head
			m1f = m1f.Next
			m2f.Next = m1f.Next
		} else {
			m2f.Next = head
			m2f = m2f.Next
		}
		head = head.Next
	}

	if m1f != nil {
		m1f.Next = m2s
	}
	if m2s == init {
		m2f.Next = nil
	}
	if m1s != nil {
		return m1s
	} else {
		return m2s
	}
}
