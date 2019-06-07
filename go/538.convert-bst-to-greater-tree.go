/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	convert(root, 0)
	return root
}

func convert(root *TreeNode, acc int) int {
	if root == nil {
		return acc
	}

	acc = convert(root.Right, acc)
	root.Val += acc
	acc = root.Val
	acc = convert(root.Left, acc)
	
	return acc
}

