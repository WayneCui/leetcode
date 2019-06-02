/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode { Val: nums[0],}
	} else if n == 2 {
		return &TreeNode {
			Val: nums[0],
			Right: &TreeNode {
				Val: nums[1],
			},
		}
	}

	mid := n / 2
	root := TreeNode{
		Val: nums[mid],
		Left: sortedArrayToBST(nums[0: mid]),
		Right: sortedArrayToBST(nums[mid + 1:]),
	}

	return &root
}

