/*
 * @lc app=leetcode id=998 lang=golang
 *
 * [998] Maximum Binary Tree II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree0(root *TreeNode, val int) *TreeNode {
    if root == nil {
		return &TreeNode {
			Val: val,
		}
	}

	if val > root.Val {
		return &TreeNode { 
			Val: val,
			Left: root,
		}
	} else {
		root.Right = insertIntoMaxTree(root.Right, val)
		return root
	}
}

func insertIntoMaxTree1(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{ Val: val, }
	if root == nil { return node }
	if val > root.Val { 
		node.Left = root
		return node
	}

	curr := root
	for curr.Right != nil && val < curr.Right.Val {
		curr = curr.Right
	}

	node.Left = curr.Right
	curr.Right = node
	return root
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	node := &TreeNode { Val: val, }

	dummy := &TreeNode{ Right: root, }

	curr := dummy
	for curr.Right != nil && val < curr.Right.Val {
		curr = curr.Right
	}
	node.Left = curr.Right
	curr.Right = node
	return dummy.Right
}

