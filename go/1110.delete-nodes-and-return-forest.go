/*
 * @lc app=leetcode id=1110 lang=golang
 *
 * [1110] Delete Nodes And Return Forest
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	to_delete_map := make(map[int]int)
	for _, v := range(to_delete) {
		to_delete_map[v] = 1
	}

	res := make([]*TreeNode, 0)
	candidates := []*TreeNode{root}
    layer := []*TreeNode{root}
	for len(layer) > 0 {
		nextLayer := []*TreeNode {}
		for _, node := range(layer) {
			if _, ok := to_delete_map[node.Val]; ok {
				if node.Left != nil {
					candidates = append(candidates, node.Left)

				}

				if node.Right != nil {
					candidates = append(candidates, node.Right)
				}
			}

			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)

				if _, ok := to_delete_map[node.Left.Val]; ok {
					node.Left = nil
				}
			}

			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)

				if _, ok := to_delete_map[node.Right.Val]; ok {
					node.Right = nil
				}
			}
		}

		layer = nextLayer
	}

	for _, node := range(candidates) {
		if _, ok := to_delete_map[node.Val]; !ok {
			res = append(res, node)
		}
	}

	return res
}

