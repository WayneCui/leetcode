/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{}
	n := len(preorder)
	if n == 0 { return nil }
	if n == 1 { 
		root.Val = preorder[0]
		return root
	}

	root.Val = preorder[0]
	root_i := rootIndexInorder(inorder, root.Val)

	if root_i > 0 {
		root.Left = buildTree(preorder[1 : root_i + 1], inorder[0 : root_i])
	}
	
	if root_i < n - 1 {
		root.Right = buildTree(preorder[root_i + 1 : ], inorder[root_i + 1 : ])
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

