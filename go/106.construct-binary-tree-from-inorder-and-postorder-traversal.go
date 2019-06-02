/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(inorder []int, postorder []int) *TreeNode {
    root := &TreeNode{}
	n := len(postorder)
	if n == 0 { return nil }
	if n == 1 { 
		root.Val = postorder[0]
		return root
	}

	root.Val = postorder[n - 1]
	root_i := rootIndexInorder(inorder, root.Val)

	if root_i > 0 {
		root.Left = buildTree(inorder[0 : root_i], postorder[0 : root_i])
	}
	
	if root_i < n - 1 {
		root.Right = buildTree(inorder[root_i + 1 : ], postorder[root_i : n - 1])
	}

	return root
}

func rootIndexInorder(inorder []int, val int) int {
	var root_i int
	n := len(inorder)
	for root_i = 0; root_i < n; root_i++ {
		if inorder[root_i] == val {
			break
		}
	}
	return root_i
}

