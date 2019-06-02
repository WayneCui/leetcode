/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	if root == nil { return }
	var memo []*TreeNode
	memo = cache(root, memo)

	r := root
	for i, node := range memo {
		if i == 0 { continue }
		r.Right = node
		r.Left = nil
		r = node
	}
}

func cache(root *TreeNode, memo []*TreeNode) []*TreeNode{
	if root == nil {
		return memo
	}
	memo = append(memo, root)
	memo = cache(root.Left, memo)
	memo = cache(root.Right, memo)
	return memo
}

// credits to https://leetcode.com/problems/flatten-binary-tree-to-linked-list/discuss/36977/My-short-post-order-traversal-Java-solution-for-share
func flatten_2(root *TreeNode)  {
	if root == nil {
		return
	}

	var next *TreeNode
	flattenHelper(root, next)
}

func flattenHelper(root *TreeNode, next *TreeNode) *TreeNode {
	if root == nil {
		return next
	}
	next = flattenHelper(root.Right, next)
	next = flattenHelper(root.Left, next)

	root.Right = next
	root.Left = nil
	next = root

	return next
}

