/*
 * @lc app=leetcode id=1008 lang=golang
 *
 * [1008] Construct Binary Search Tree from Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 { return nil }
	if n == 1 { return &TreeNode{ Val: preorder[0], }}

	root := &TreeNode { Val: preorder [0], }
	idx := search(preorder[1:], preorder[0])
	root.Left = bstFromPreorder(preorder[1: idx + 1])
	root.Right = bstFromPreorder(preorder[idx + 1:])
	return root
}

func search(nums []int, target int) int {
	for i, v := range(nums) {
		if v > target {
			return i
		}
	}

	return len(nums)
}

//another way using binary search
func bstFromPreorder(preorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 { return nil }
	if n == 1 { return &TreeNode{ Val: preorder[0], }}

	root := &TreeNode { Val: preorder [0], }
	idx := binarySearch(preorder[1:], preorder[0])
    idx = -idx - 1
	root.Left = bstFromPreorder(preorder[1: idx + 1])
	root.Right = bstFromPreorder(preorder[idx + 1:])
	return root
}

func binarySearch(nums []int, target int) int {
	n := len(nums)

	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -(low + 1)
}

