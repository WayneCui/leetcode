/*
 * @lc app=leetcode id=783 lang=golang
 *
 * [783] Minimum Distance Between BST Nodes
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	vals := inorderTraversal(root)
	n := len(vals)
	if n == 0 || n == 1 {
		return 0
	}

	min := math.MaxInt64
	for i := 0; i < n - 1; i++ {
		df := diff(vals[i], vals[i + 1])
		if df < min {
			min = df
		}
	}

	return min
}

func diff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func inorderTraversal(root *TreeNode) []int {
    if root == nil { 
		return []int{}
	}

	output := []int{}
	if root.Left != nil {
		output = append(output, inorderTraversal(root.Left)...)
	}

	output = append(output, root.Val)

	if root.Right != nil {
		output = append(output, inorderTraversal(root.Right)...)
	}

	return output
}

