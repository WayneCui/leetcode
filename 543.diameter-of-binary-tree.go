/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	_, diam := depth(root, 0)
	return diam
}

func depth(root *TreeNode, maxDiam int) (dp int, diam int) {
	if root == nil { return 0, maxDiam }
	l, ld := depth(root.Left, maxDiam)
	r, rd := depth(root.Right, maxDiam)

	return max(l, r) + 1, max(ld, rd, l + r, maxDiam)
}

func max(nums ...int) int {
	m := 0
	for _, num := range(nums) {
		if m < num {
			m = num
		}
	}

	return m
}

