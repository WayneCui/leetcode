/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil || 
		(root.Left == nil && root.Right == nil) {
		return true
	} 

	return isSym(root.Left, root.Right)
}

func isSym(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil { return true }
	if root1 == nil || root2 == nil { return false }
	if root1.Val != root2.Val { return false }
	return isSym(root1.Left, root2.Right) && isSym(root1.Right, root2.Left)
}

