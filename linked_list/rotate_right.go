package main

// O(n)
// https://leetcode.com/problems/rotate-list
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n := 1
	current := head
	for current.Next != nil {
		current = current.Next
		n++
	}
	current.Next = head
	if k >= n {
		k %= n
	}
	leftRotations := n - k
	prevHead := head
	for range leftRotations {
		prevHead = head
		head = head.Next
	}
	prevHead.Next = nil
	return head
}
