/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := collectLeaf(root1)
	leaves2 := collectLeaf(root2)

	if len(leaves1) != len(leaves2) {
		return false
	} else {
		n := len(leaves1)
		for i := 0; i < n; i++ {
			if leaves1[i] != leaves2[i] {
				return false
			}
		}

		return true
	}
}

func collectLeaf(root *TreeNode) []int {
	output := []int{}
	if isLeaf(root) {
		output = append(output, root.Val)
	} else {
		if root.Left != nil {
			output = append(output, collectLeaf(root.Left)...)
		}

		if root.Right != nil {
			output = append(output, collectLeaf(root.Right)...)
		}
	}

	return output
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

