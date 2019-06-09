/*
 * @lc app=leetcode id=993 lang=golang
 *
 * [993] Cousins in Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil || root.Val == x || root.Val == y {
		return false
	}

	layer := []*TreeNode { root }
	for {
		if len(layer) == 0 { break }

		nextLayer := []*TreeNode {}
		var states []int

		for _, node := range(layer) {
			state := 0

			if node.Left != nil {
				if node.Left.Val == x || node.Left.Val == y {
					state += 1
				}
				nextLayer = append(nextLayer, node.Left)
			}

			if node.Right != nil {
				if node.Right.Val == x || node.Right.Val == y {
					state += 1
				}
				nextLayer = append(nextLayer, node.Right)
			}

			//same parent
			if state == 2 {
				return false
			}
			states = append(states, state)
		}

		count := 0
		for _, num := range(states) {
			if num == 1 {
				count += 1
			}
		}

		if count == 1 {
			return false
		} else if count == 2 {
			return true
		}

		layer = nextLayer
	}

	return false
}

