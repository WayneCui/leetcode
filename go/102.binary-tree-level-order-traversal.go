/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := [][]*TreeNode{ {root} }
	flag := 0
	for len(queue) > flag {
		layer := queue[flag]
		newLayer := []*TreeNode{}
		layerVal := []int{}
		for _, node := range layer {
			if node != nil {
				layerVal = append(layerVal, node.Val)
				if node.Left != nil {
					newLayer = append(newLayer, node.Left)
				}
				
				if node.Right != nil {
					newLayer = append(newLayer, node.Right)
				}
			}
		}

		result = append(result, layerVal)
		if len(newLayer) > 0 {
			queue = append(queue, newLayer)
		}
		flag += 1
	}

	return result
}