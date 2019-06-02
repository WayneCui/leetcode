/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil { return 0 }
	return pathSumFrom(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathSumFrom(node *TreeNode, sum int) int {
	if node == nil { return 0 }
	count := 0
	if node.Val == sum {
		count += 1 
	}

	return count + pathSumFrom(node.Left, sum - node.Val) + pathSumFrom(node.Right, sum - node.Val)
}

