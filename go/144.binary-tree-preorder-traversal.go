/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//recursive
func preorderTraversal0(root *TreeNode) []int {
	var output []int
	return pTraversal(root, output)
}

func pTraversal(root *TreeNode, output []int) []int {
	if root == nil { return output }
	output = append(output, root.Val)
	if root.Left != nil {
		output = pTraversal(root.Left, output)
	}

	if root.Right != nil {
		output = pTraversal(root.Right, output)
	}

	return output
}

//iterative
func preorderTraversal(root *TreeNode) []int {
	output := []int{}
	if root == nil {
		return output
	}

	stack := []*TreeNode { root }
	var node *TreeNode
	for len(stack) > 0 {
		node, stack = pop(stack)
		if node.Right != nil {
			stack = push(stack, node.Right)
		}
		
		if node.Left != nil {
			stack = push(stack, node.Left)
		}

		output = append(output, node.Val)
	} 

	return output
}

func pop(a []*TreeNode) (poped *TreeNode, remained []*TreeNode) {
	x, a := a[len(a)-1], a[:len(a)-1]
	return x, a
}

func push(a []*TreeNode, node *TreeNode) []*TreeNode {
	return append(a, node)
}

