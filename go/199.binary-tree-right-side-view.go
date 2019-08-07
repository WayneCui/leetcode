/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil { return []int{} }
	
	res := []int {}
	layer := []*TreeNode{ root }
	for len(layer) > 0 {
		nextLayer := []*TreeNode {}
		theone := 0
		for _, node := range(layer) {
			theone = node.Val

			if node.Left != nil { nextLayer = append(nextLayer, node.Left) }
			if node.Right != nil { nextLayer = append(nextLayer, node.Right)}
		}

		res = append(res, theone)
		layer = nextLayer
	}

	return res
}

