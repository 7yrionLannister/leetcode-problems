package main

import "fmt"

func main() {
	// tree := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}
	tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	flatten(tree)
	fmt.Println("chao")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/flatten-binary-tree-to-linked-list
// O(N) time and O(1) space
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	badRoot := root
	for badRoot.Left == nil && badRoot.Right != nil {
		badRoot = badRoot.Right
	}
	if badRoot.Left == nil && badRoot.Right == nil {
		// it's the end, not really a bad root
		return
	}
	badRootsLastRightNode := badRoot.Left
	for badRootsLastRightNode.Right != nil {
		badRootsLastRightNode = badRootsLastRightNode.Right
	}
	badRootsLastRightNode.Right = badRoot.Right
	badRoot.Right = badRoot.Left
	badRoot.Left = nil
	if badRoot.Right != nil {
		flatten(badRoot.Right)
	}
}
