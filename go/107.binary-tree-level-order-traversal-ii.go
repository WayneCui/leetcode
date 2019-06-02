/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrderBottom1(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil { return result }
	
	layers := [][]*TreeNode{{root}}
	flag := 0
	for flag < len(layers) {
		layer := layers[flag]
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

        if len(newLayer) > 0 {
			layers = append(layers, newLayer)
		}

        result = insert(result, layerVal)
		flag += 1
		
	}

	return result

}

func insert(s [][]int, v []int) [][]int {
	s = append(s, []int{})
	copy(s[1:], s[0:])
	s[0] = v
    return s
}

func levelOrderBottom(root *TreeNode) [][]int {
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

	for i := len(result)/2-1; i >= 0; i-- {
		opp := len(result)-1-i
		result[i], result[opp] = result[opp], result[i]
	}

	return result
}
