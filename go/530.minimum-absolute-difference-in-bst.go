/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	vals := inorder(root)
	
	min := math.MaxInt64
	for i := 0; i < len(vals) - 1; i++ {
		diff := vals[i + 1] - vals[i]
		if diff < min {
			min = diff
		}
	}

	return min
}

func inorder(root *TreeNode) []int {
	output := []int{}
	if root == nil { return output }
	output = append(output, inorder(root.Left)...)
	output = append(output, root.Val)
	output = append(output, inorder(root.Right)...)

	return output
}

