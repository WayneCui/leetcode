/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    if root == nil { 
		return []int{}
	}

	output := []int{}
	if root.Left != nil {
		output = append(output, inorderTraversal(root.Left)...)
	}

	output = append(output, root.Val)

	if root.Right != nil {
		output = append(output, inorderTraversal(root.Right)...)
	}

	return output
}

//iteratively
func inorderTraversal1(root *TreeNode) []int {
	var output []int
	if root == nil {
		return output
	}
	
	stack := []*TreeNode{ root }
	visited := []*TreeNode{}
	for len(stack) > 0 {
		node := peek(stack)

		if node.Left != nil && !find(visited, node.Left) {
			stack = push(stack, node.Left)
		} else {
			output = append(output, node.Val)
			node, stack = pop(stack)
			visited = append(visited, node)

			if node.Right != nil {
				stack = push(stack, node.Right)
			}
		}
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

func peek(a []*TreeNode) *TreeNode {
	return a[len(a) - 1]
}

func find(nodes []*TreeNode, target *TreeNode) bool {
	n := len(nodes)
	if n <= 0 { return false }
	for i := n - 1; i >= 0; i-- {
		if nodes[i] == target {
			return true
		}
	}

	return false
}

