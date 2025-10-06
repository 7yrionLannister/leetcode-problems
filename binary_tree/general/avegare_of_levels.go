package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	averageOfLevels(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}})
}

// https://leetcode.com/problems/average-of-levels-in-binary-tree
var (
	levelsSum    []float64
	nodesInLevel []int
)

func averageOfLevels(root *TreeNode) []float64 {
	levelsSum = make([]float64, 0)
	nodesInLevel = make([]int, 0)
	averageOfLevelsRecursive(root, 0)
	for i := range levelsSum {
		levelsSum[i] /= float64(nodesInLevel[i])
	}
	return levelsSum
}

func averageOfLevelsRecursive(root *TreeNode, i int) {
	if root != nil {
		if i >= len(levelsSum) {
			levelsSum = append(levelsSum, 0)
			nodesInLevel = append(nodesInLevel, 0)
		}
		levelsSum[i] += float64(root.Val)
		nodesInLevel[i]++
		averageOfLevelsRecursive(root.Left, i+1)
		averageOfLevelsRecursive(root.Right, i+1)
	}
}
