/*
 * @lc app=leetcode id=513 lang=golang
 *
 * [513] Find Bottom Left Tree Value
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	layer := []*TreeNode { root }
	var leftMost int
	for len(layer) > 0 {
		nextLayer := []*TreeNode {}

		for _, node := range(layer) {
			if node.Left != nil { nextLayer = append(nextLayer, node.Left)}
			if node.Right != nil { nextLayer = append(nextLayer, node.Right)}
		}

		leftMost = layer[0].Val
		layer = nextLayer
	}

	return leftMost
}

