package main

// https://leetcode.com/problems/flatten-binary-tree-to-linked-list
// O(N) time and O(N) space

var nodes []*TreeNode

func flatten(root *TreeNode) {
	nodes = make([]*TreeNode, 0)
	preorderTraversal(root)
	n := len(nodes)
	for i, node := range nodes {
		node.Left = nil
		if i < n-1 {
			node.Right = nodes[i+1]
		}
	}
}

func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	nodes = append(nodes, root)
	if root.Left != nil {
		preorderTraversal(root.Left)
	}
	if root.Right != nil {
		preorderTraversal(root.Right)
	}
}
