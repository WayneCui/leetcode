/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
	if root.Left == nil && root.Right == nil {
		return 1
	}
	
	left := math.MinInt32
	if root.Left != nil {
		left = maxDepth(root.Left) + 1
	}
	
	right := math.MinInt32
	if root.Right != nil {
		right = maxDepth(root.Right) + 1
	}

	if left > right {
		return left 
	} else {
		return right
	}
}

