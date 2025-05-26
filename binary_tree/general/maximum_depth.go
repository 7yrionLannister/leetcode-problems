package general

// Time: O(n)
func maxDepth(root *TreeNode) int {
	return maxDepthRecursive(root, 0)
}

func maxDepthRecursive(root *TreeNode, currentDepth int) int {
	if root == nil {
		return currentDepth
	}
	currentDepth++
	leftdepth := maxDepthRecursive(root.Left, currentDepth)
	rightDepth := maxDepthRecursive(root.Right, currentDepth)
	return max(currentDepth, leftdepth, rightDepth)
}
