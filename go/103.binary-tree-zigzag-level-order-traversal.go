/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var output [][]int
	if root == nil { return output }

	level := []*TreeNode{ root }
	flag := 1
	for len(level) > 0 {
		n := len(level)
		values := make([]int, n)
		nextLevel := []*TreeNode {}
		
		for index, node := range(level) {
			if flag == 1 {
				values[index] = node.Val
			} else {
				values[n - index - 1] = node.Val
			}
			
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		flag = -flag
		output = append(output, values)
		level = nextLevel
	}

	return output
}

