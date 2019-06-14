/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	var output []float64
	level := []*TreeNode{ root }
	for {
		if len(level) <= 0 { break }

		nextLevel := []*TreeNode {}
		sum := 0
		for _, node := range(level) {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		avg := float64(sum) / float64(len(level))
		output = append(output, avg)
		level = nextLevel
	}

	return output
}

