/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes0(root *TreeNode) int {
	if root == nil { return 0 }
	
	layer := []*TreeNode { root }
	count := 0
	for len(layer) != 0 {
		count += len(layer)
		nextLayer := []*TreeNode {}
		for _, node := range(layer) {
			if node.Left != nil { nextLayer = append(nextLayer, node.Left) }
			if node.Right != nil { nextLayer = append(nextLayer, node.Right)}
		}

		layer = nextLayer
	}

	return count
}

func countNodes(root *TreeNode) int {
	if root == nil { return 0 }
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

