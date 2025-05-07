package main

import "fmt"

func main() {
	n2 := &Node{2, nil, nil}
	n2.Random = n2
	n1 := &Node{1, n2, n2}
	res := copyRandomList(n1)
	fmt.Println(res)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode.com/problems/copy-list-with-random-pointer
// Time: O(n)
// Space: O(2n) == O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := &Node{
		Val: head.Val,
	}
	ref := newHead
	pointerToPointer := make(map[*Node]*Node)
	pointerToPointer[head] = newHead
	for head != nil {
		if head.Next != nil {
			nextPointer, present := pointerToPointer[head.Next]
			if present {
				ref.Next = nextPointer
			} else {
				ref.Next = &Node{
					Val: head.Next.Val,
				}
				pointerToPointer[head.Next] = ref.Next
			}
		}
		if head.Random != nil {
			randomPointer, present := pointerToPointer[head.Random]
			if present {
				ref.Random = randomPointer
			} else {
				ref.Random = &Node{
					Val: head.Random.Val,
				}
				pointerToPointer[head.Random] = ref.Random
			}
		}
		head = head.Next
		ref = ref.Next
	}
	return newHead
}
