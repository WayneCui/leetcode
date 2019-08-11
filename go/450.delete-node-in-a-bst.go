/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
    if root.Val == key {
		root = rebalance(root)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func rebalance(toRemove *TreeNode) *TreeNode {
	if toRemove.Left == nil {
		return toRemove.Right
	}
	if toRemove.Right == nil {
		return toRemove.Left
	}
	
	prev := toRemove
	node := toRemove.Right
	for node != nil {
		prev = node
		node = node.Left
	}
	prev.Left = toRemove.Left
	return toRemove.Right
}

