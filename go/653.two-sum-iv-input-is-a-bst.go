/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	memo := make(map[int]bool)
	dsf(root, memo)

	for key, _ := range(memo) {
		if key == k - key {
			continue
		}
		_, ok := memo[k - key]
		if ok {
			return true
		}
	}

	return false
}

func dsf(root *TreeNode, memo map[int]bool) {
	if root == nil {
		return 
	}
	memo[root.Val] = true
	dsf(root.Left, memo)
	dsf(root.Right, memo)
}

