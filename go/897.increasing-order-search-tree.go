/*
 * @lc app=leetcode id=897 lang=golang
 *
 * [897] Increasing Order Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	stack := []*TreeNode { root }
	
	superRoot := &TreeNode {}
	curr := superRoot
	for len(stack) > 0 {
		node := peek(&stack)

		if node.Left == nil {
			curr.Right = node
			curr = curr.Right

			pop(&stack)
			if node.Right != nil {
				push(&stack, node.Right)
			}
		} else {
			push(&stack, node.Left)
			node.Left = nil
		}
	}

	return superRoot.Right
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

func push(stack *[]*TreeNode, node *TreeNode) *[]*TreeNode {
	return append(*stack, node)
}

func pop(stack *[]*TreeNode) *TreeNode {
	n := len(stack)
	if n > 0 {
		return stack[ n - 1]
		stack = stack[0 : n - 1]
	}
	
	return nil
}

func peek(stack *[]*TreeNode) *TreeNode {
	n := len(stack)
	if n > 0 {
		return stack[ n - 1]
	}
	
	return nil
}

