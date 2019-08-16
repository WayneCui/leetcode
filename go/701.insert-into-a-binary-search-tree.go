/*
 * @lc app=leetcode id=701 lang=golang
 *
 * [701] Insert into a Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode {
				Val: val,
			}
		} else {
			insertIntoBST(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode {
				Val: val,
			}
		} else {
			insertIntoBST(root.Left, val)
		}
	}

	return root
}

