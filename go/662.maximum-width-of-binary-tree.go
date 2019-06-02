/*
 * @lc app=leetcode id=662 lang=golang
 *
 * [662] Maximum Width of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree0(root *TreeNode) int {
	if root == nil { return 0 }

	nodes := []*TreeNode{ root }
	maxWidth := 0
	for {
		n := len(nodes)
		if n == 0 { 
			break
		} else if maxWidth < len(nodes) { 
			maxWidth = len(nodes) 
		}

		nodes = nextLevel0(nodes)
	}
	return maxWidth
}

func nextLevel0( nodes []*TreeNode) []*TreeNode {
	if len(nodes) == 0 {
		return []*TreeNode{}
	}

	var newNodes []*TreeNode
	for _, node := range(nodes) {
		if node == nil {
			newNodes = append(newNodes, nil)
			newNodes = append(newNodes, nil)
		} else {
			newNodes = append(newNodes, node.Left)
			newNodes = append(newNodes, node.Right)
		}
	}
	newNodes = trim(newNodes)

	return newNodes
}

func trim(nodes []*TreeNode)[]*TreeNode {
	n := len(nodes)
	st := 0
	for st < n && nodes[st] == nil {
		st += 1
	}

	for st < n && nodes[n - 1] == nil {
		n -= 1
	}

	return nodes[st : n ]
}

func max(nums ...int) int {
	m := 0
	for _, num := range(nums) {
		if m < num {
			m = num
		}
	}

	return m
}


func widthOfBinaryTree(root *TreeNode) int {
	if root == nil { return 0 }

	nodes := []*TreeNode{ root }
	idx := []int{1}
	maxWidth := 0
	for len(nodes) > 0{
		n := len(nodes)
		if maxWidth < (idx[n - 1] - idx[0] + 1) { 
			maxWidth = idx[n - 1] - idx[0] + 1
		} 

		nodes, idx = nextLevel(nodes, idx)
	}
	return maxWidth
}

func nextLevel( nodes []*TreeNode, idx []int) ([]*TreeNode, []int)  {
	var nextNodes []*TreeNode
	var nextIdx []int
	for i, node := range(nodes) {
		if node.Left != nil {
			nextNodes = append(nextNodes, node.Left)
			nextIdx = append(nextIdx, 2 * idx[i])
		}
		
		if node.Right != nil {
			nextNodes = append(nextNodes, node.Right)
			nextIdx = append(nextIdx, 2 * idx[i] + 1)
		}
		
	}

	return nextNodes, nextIdx
}


