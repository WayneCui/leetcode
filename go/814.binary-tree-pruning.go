/*
 * @lc app=leetcode id=814 lang=golang
 *
 * [814] Binary Tree Pruning
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	//postorder
	pruneTree(root.Left)
	pruneTree(root.Right)
    
	if isLeaf(root.Left) && root.Left.Val == 0 {
		root.Left = nil
	}
	if isLeaf(root.Right) && root.Right.Val == 0 {
		root.Right = nil
	}
	
	if isLeaf(root) && root.Val == 0 {
		return nil
	} else {
		return root
	}
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

// a more concise way
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil { return nil }

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if isLeaf(root) && root.Val == 0 { 
		return nil 
	} else {
		return root
	}
}
