package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n)
// Space: O(1)
// https://leetcode.com/problems/linked-list-cycle
func hasCycle(head *ListNode) bool {
	reference := new(ListNode)
	for head != nil {
		if head.Next == reference {
			return true
		}
		prevHead := head
		head = head.Next
		prevHead.Next = reference
	}
	return false
}
