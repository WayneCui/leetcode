/*
 * @lc app=leetcode id=1123 lang=golang
 *
 * [1123] Lowest Common Ancestor of Deepest Leaves
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := postOrder(root)
	return lca
}

func postOrder(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	left, lnode := postOrder(root.Left)
	right, rnode := postOrder(root.Right)

	var lca *TreeNode
	if left > right {
		lca = lnode
	} else if right > left {
		lca = rnode
	} else {
		lca = root
	}

	return max(left, right) + 1, lca
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
