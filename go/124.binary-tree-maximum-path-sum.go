/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxPathSum(root *TreeNode) int {
	if root == nil {
		return math.MinInt32
	}
	_, globalMax := dfs(root, math.MinInt64)
	return globalMax
}

func dfs(root *TreeNode, globalMax int) (localMax int, newGlobalMax int) {
	var lmx1, lmx2 int
	gmx1 := math.MinInt32
	gmx2 := math.MinInt32
	if root.Left != nil {
		lmx1, gmx1 = dfs(root.Left, globalMax)
	}

	if root.Right != nil {
		lmx2, gmx2 = dfs(root.Right, globalMax)
	}

	lmx := max(root.Val, lmx1 + root.Val, lmx2 + root.Val)
	gmx := max(root.Val, gmx1, gmx2, lmx, lmx1 + lmx2 + root.Val, globalMax)

	return lmx, gmx
}

func max(args ...int) int {
	max := math.MinInt64
	for _, num := range(args) {
		if num > max {
			max = num
		}
	}

	return max
}

