/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// recursive
func postorderTraversal0(root *TreeNode) []int {
	var output []int
	return postTraversal(root, output)
}

func postTraversal(root *TreeNode, output []int) []int {
	if root == nil { return output }
	if root.Left != nil {
		output = postTraversal(root.Left, output)
	}

	if root.Right != nil {
		output = postTraversal(root.Right, output)
	}

	output = append(output, root.Val)

	return output
}

//iterative
func postorderTraversal(root *TreeNode) []int {
	output := []int {}
	if root == nil { return output }

	stack := []*TreeNode { root }
	for len(stack) > 0 {
		node := peek(stack)

		if isLeaf(node) {
			output = append(output, node.Val)
			node, stack = pop(stack)
		} else {
			if node.Right != nil {
				stack = push(stack, node.Right)
				node.Right = nil
			}
	
			if node.Left != nil {
				stack = push(stack, node.Left)
				node.Left = nil
			}
		}
	}

	return output
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

func pop(a []*TreeNode) (poped *TreeNode, remained []*TreeNode) {
	x, a := a[len(a)-1], a[:len(a)-1]
	return x, a
}

func push(a []*TreeNode, node *TreeNode) []*TreeNode {
	return append(a, node)
}

func peek(a []*TreeNode) *TreeNode {
	return a[len(a) - 1]
}

