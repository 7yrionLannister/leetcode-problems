package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(hasPathSum(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 5))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumRecursive(root, targetSum, 0)
}

func hasPathSumRecursive(root *TreeNode, targetSum int, sum int) bool {
	if root != nil {
		sum += root.Val
		if root.Left == nil && root.Right == nil {
			// leaf
			if sum == targetSum {
				return true
			}
		}
		return hasPathSumRecursive(root.Left, targetSum, sum) || hasPathSumRecursive(root.Right, targetSum, sum)
	}
	return false
}
