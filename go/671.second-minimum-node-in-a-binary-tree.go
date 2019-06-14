/*
 * @lc app=leetcode id=671 lang=golang
 *
 * [671] Second Minimum Node In a Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || isLeaf(root) {
		return -1
	}
	if root.Val != root.Left.Val {
		return minPosInt(root.Left.Val, findSecondMinimumValue(root.Right))
	} else if root.Val != root.Right.Val {
		return minPosInt(root.Right.Val, findSecondMinimumValue(root.Left))
	} else {
		return minPosInt(findSecondMinimumValue(root.Left), findSecondMinimumValue(root.Right))
	}
}

func minPosInt(a, b int) int {
	if a < 0 && b < 0 {
		return -1
	} else if a < 0 {
		return b
	} else if b < 0 {
		return a
	} else if a < b {
			return a
	} else {
		return b
	}
}

func isLeaf(root *TreeNode) bool {
	return root != nil && root.Left == nil && root.Right == nil
}

