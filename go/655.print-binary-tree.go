/*
 * @lc app=leetcode id=655 lang=golang
 *
 * [655] Print Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	memo := [][]string{}
	layer := []*TreeNode { root }
	for len(layer) > 0 {
		nextLayer := []*TreeNode {}
		layerVal := []string{}
		nilCount := 0
		for _, node := range(layer) {
			if node == nil {
				layerVal = append(layerVal, "")
				// nextLayer = append(nextLayer, node, node)
				nextLayer = append(nextLayer, nil, nil)
				nilCount += 2
			} else {
				layerVal = append(layerVal, strconv.Itoa(node.Val))
				nextLayer = append(nextLayer, node.Left)
				nextLayer = append(nextLayer, node.Right)
				if node.Left == nil { nilCount += 1}
				if node.Right == nil { nilCount += 1}
			}
		}

		memo = append(memo, layerVal)

		if len(nextLayer) == nilCount {
			break
		}
		layer = nextLayer
	}

	// fmt.Println( memo )

	h := len(memo)
	res := make([][]string, h)
	p := 0
	for i := h - 1; i >= 0; i-- {
		layer := []string{}
		for j, v := range(memo[i]) {
			layer = append(layer, padding(v, p)...)
			if j < len(memo[i]) - 1 {
				layer = append(layer, "")
			}
		}

		res[i] = layer
		p = p * 2 + 1
	}

	return res
}

func padding(v string, c int) []string {
	res := []string{}
	res = append(res, make([]string, c)...)
	res = append(res, v)
	res = append(res, make([]string, c)...)

	return res
}

