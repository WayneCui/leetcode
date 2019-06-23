/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	memo := []int{}
	inorder(root, &memo, k)
	if len(memo) >= k {
		return memo[k - 1]
	} else {
		return math.MinInt32
	}
}

func inorder(root *TreeNode, memo *[]int, k int) {
	if len(*memo) < k {
		if root.Left != nil {
			inorder(root.Left, memo, k)
		}

		*memo = append(*memo, root.Val)

		if root.Right != nil {
			inorder(root.Right, memo, k)
		}
	}
}
