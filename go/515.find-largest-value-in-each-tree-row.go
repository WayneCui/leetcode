/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil { return []int{} }
	layer := []*TreeNode { root }
	res := []int{}
	for len(layer) > 0 {
		nextLayer := []*TreeNode {}
		largest := math.MinInt64
		for _, node := range(layer) {
			//core work
			if node.Val > largest {
				largest = node.Val
			}

			//iterative
			if node.Left != nil { nextLayer = append(nextLayer, node.Left)}
			if node.Right != nil { nextLayer = append(nextLayer, node.Right)}
		}

		res = append(res, largest)
		layer = nextLayer
	}


	return res
}

