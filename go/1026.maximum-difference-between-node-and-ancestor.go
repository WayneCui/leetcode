/*
 * @lc app=leetcode id=1026 lang=golang
 *
 * [1026] Maximum Difference Between Node and Ancestor
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxAncestorDiff(root *TreeNode) int {
	_, _, diff := dfs(root)
	return diff
}

func dfs(root *TreeNode) (mx int, mn int, diff int) {
	if isLeaf(root) {
		return root.Val, root.Val, 0
	}

	lx, ln, ldiff := math.MinInt32, math.MaxInt32, 0
	if root.Left != nil {
		lx, ln, ldiff = dfs(root.Left)
	}
	rx, rn, rdiff := math.MinInt32, math.MaxInt32, 0
	if root.Right != nil {
		rx, rn, rdiff = dfs(root.Right)
	}

	mmx := max(lx, rx, root.Val)
	mmn := min(ln, rn, root.Val)
	df := max(ldiff, rdiff, abs(root.Val - max(lx, rx)), abs(root.Val - min(ln, rn))) 

	return mmx, mmn, df
}

func max(a...int) int {
	max := math.MinInt32
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	return max
}

func min(a...int) int {
	min := math.MaxInt32
	for _, v := range a {
		if v < min {
			min = v
		}
	}

	return min
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func isLeaf(root *TreeNode) bool {
	return root != nil && root.Left == nil && root.Right == nil
}
