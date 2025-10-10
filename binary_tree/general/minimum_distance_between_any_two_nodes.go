package main

import "math"

func main() {
	minDiff := getMinimumDifference(&TreeNode{1, nil, &TreeNode{3, &TreeNode{2, nil, nil}, nil}})
	println(minDiff)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/minimum-absolute-difference-in-bst
// O(n) time
func getMinimumDifference(root *TreeNode) int {
	// as the number of nodes in the tree ranges from [2..10^4], no need to verify root != nil
	lm, rm, minDiff := getMinimumDifferenceAndLeftMostAndRightMostValues(root, math.MaxInt)
	if lm != nil && *lm != root.Val {
		ld := (root.Val - *lm)
		if ld < 0 {
			ld *= -1
		}
		minDiff = min(minDiff, ld)
	}
	if rm != nil && *rm != root.Val {
		rd := (root.Val - *rm)
		if rd < 0 {
			rd *= -1
		}
		minDiff = min(minDiff, rd)
	}
	return minDiff
}

// If we have the minimun distance for this recursive call on root, we can compare it to the previous minimum distance.
// Having the left and right most values helps us compare the relevant nodes in the left and right subtrees to compute even
// smaller differences.
func getMinimumDifferenceAndLeftMostAndRightMostValues(root *TreeNode, currentMinimumDifference int) (leftMostVal, rightMostVal *int, newMinimumDifference int) {
	if currentMinimumDifference == 1 {
		return nil, nil, 1
	}
	rightMostVal = &root.Val
	leftMostVal = &root.Val
	if root.Left == nil && root.Right == nil { // leaf, does not change the diff
		return leftMostVal, rightMostVal, currentMinimumDifference
	}
	newMinimumDifference = currentMinimumDifference
	if root.Left != nil {
		leftMostVal = &root.Left.Val
		// recursive left call
		lm, rm, nd := getMinimumDifferenceAndLeftMostAndRightMostValues(root.Left, newMinimumDifference)
		rd := math.MaxInt
		if rm != nil {
			rd = (root.Val - *rm)
			if rd < 0 {
				rd *= -1
			}
		}
		newMinimumDifference = min(newMinimumDifference, nd, rd)
		if lm != nil {
			leftMostVal = lm
		}
	}
	if root.Right != nil {
		// recursive right call
		rightMostVal = &root.Right.Val
		lm, rm, nd := getMinimumDifferenceAndLeftMostAndRightMostValues(root.Right, newMinimumDifference)
		ld := math.MaxInt
		if lm != nil {
			ld = (root.Val - *lm)
			if ld < 0 {
				ld *= -1
			}
		}
		newMinimumDifference = min(newMinimumDifference, nd, ld)
		if rm != nil {
			rightMostVal = rm
		}
	}
	return leftMostVal, rightMostVal, newMinimumDifference
}
