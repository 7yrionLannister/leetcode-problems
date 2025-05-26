package main

func invertTree(root *TreeNode) *TreeNode {
	invertTreeRecursive(root)
	return root
}

func invertTreeRecursive(root *TreeNode) {
	if root != nil {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp
		invertTreeRecursive(root.Left)
		invertTreeRecursive(root.Right)
	}
}
