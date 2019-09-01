/*
 * @lc app=leetcode id=988 lang=golang
 *
 * [988] Smallest String Starting From Leaf
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	return helper(root, "")
}

func helper(root *TreeNode, init string) string {
	if root == nil { return "|" } // "|" > "z"
	if isLeaf(root) { return repend(init, string(root.Val + 97))}

	init = repend(init, string(root.Val + 97))
	return min(helper(root.Left, init), helper(root.Right, init))
}

func isLeaf(root *TreeNode) bool {
	return root != nil && root.Left == nil && root.Right == nil
}

func min(a, b string) string {
	if a < b {
		return a
	} else {
		return b
	}
}

func repend(init, s string) string {
	return s + init
}

