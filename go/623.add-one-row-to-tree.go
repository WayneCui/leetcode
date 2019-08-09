/*
 * @lc app=leetcode id=623 lang=golang
 *
 * [623] Add One Row to Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode {
			Val: v,
			Left: root,
		}
	}
	
	layer := []*TreeNode { root }
	depth := 0
	for len(layer) > 0 {
		depth += 1
		nextLayer := []*TreeNode {}
		for _, node := range(layer) {
			if depth == d - 1 {
				lf := node.Left
				node.Left = &TreeNode{
					Val: v,
					Left: lf,
				}

				rt := node.Right
				node.Right = &TreeNode {
					Val: v,
					Right: rt,
				}
			} else {
				if node.Left != nil { nextLayer = append(nextLayer, node.Left)  }
				if node.Right != nil { nextLayer = append(nextLayer, node.Right) }
			}
		}
		
		layer = nextLayer
	}

	return root
}

