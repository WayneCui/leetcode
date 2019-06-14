/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	_, longest := traversal(root, 0)
	
	return longest
}

func traversal(node *TreeNode, longest int) (int, int) {
	if node == nil {
		return 0, longest
	}

	lf, llf := traversal(node.Left, longest)
	rg, lrg := traversal(node.Right, longest)

	left := 0
	right := 0
	if node.Left != nil && node.Left.Val == node.Val {
		left = lf + 1
	}

	if node.Right != nil && node.Right.Val == node.Val {
		right = rg + 1
	}

	longest = max(longest, left + right, llf, lrg)

	return max(left, right), longest
}

func max(nums ...int) int {
	result := math.MinInt64
	for _, num := range(nums) {
		if num > result {
			result = num
		}
	}

	return result
}

