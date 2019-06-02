/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil { return 0 }
	if root.Left == nil && root.Right == nil {
		return 1
	}
	
	left := math.MaxInt32
	if root.Left != nil {
		left = minDepth(root.Left) + 1
	}
	
	right := math.MaxInt32
	if root.Right != nil {
		right = minDepth(root.Right) + 1
	}

	if left < right {
		return left 
	} else {
		return right
	}
}

