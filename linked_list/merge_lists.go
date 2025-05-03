package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Time: O(n)
// Space: O(1)
// https://leetcode.com/problems/merge-two-sorted-lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newHead *ListNode
	if list1 == nil {
		newHead = list2
		if list2 != nil {
			list2 = list2.Next
		}
	} else if list2 == nil {
		newHead = list1
		list1 = list1.Next
	} else if list1.Val < list2.Val {
		newHead = list1
		list1 = list1.Next
	} else {
		newHead = list2
		list2 = list2.Next
	}
	refHead := newHead
	for {
		if list1 == nil {
			if refHead != nil {
				refHead.Next = list2
			}
			break
		}
		if list2 == nil && refHead != nil {
			if refHead != nil {
				refHead.Next = list1
			}
			break
		}
		if list1.Val < list2.Val {
			refHead.Next = list1
			refHead = refHead.Next
			list1 = list1.Next
		} else {
			refHead.Next = list2
			refHead = refHead.Next
			list2 = list2.Next
		}
	}
	return newHead
}
