/*
 * @lc app=leetcode id=654 lang=golang
 *
 * [654] Maximum Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 { return nil }
	if len(nums) == 1 { 
		return &TreeNode{
			Val: nums[0],
		} 
	}

	idx, mx := max(nums)
	return &TreeNode {
		Val: mx,
		Left: constructMaximumBinaryTree(nums[0:idx]),
		Right: constructMaximumBinaryTree(nums[idx + 1:]),
	}
}

func max(nums []int) (index, val int) {
	m := 0
	idx := -1
	for i, v := range(nums) {
		if m < v { 
			m = v
			idx = i
		}
	}

	return idx, m
}

