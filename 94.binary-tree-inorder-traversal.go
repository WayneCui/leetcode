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
func inorderTraversal0(root *TreeNode) []int {
	var output []int
	return traversal(root, output)
}

func traversal(root *TreeNode, container []int) []int {
	if root == nil { return container }

	if root.Left != nil {
		container = traversal(root.Left, container)
	}
	container = append(container, root.Val)
	if root.Right != nil {
		container = traversal(root.Right, container)
	}

	return container
}

//iteratively
func inorderTraversal(root *TreeNode) []int {
	var output []int
	if root == nil {
		return output
	}
	
	queue := []*TreeNode{ root }
	visited := []*TreeNode{}
	for len(queue) > 0 {
		node := peek(queue)

		if node.Left != nil && !find(visited, node.Left) {
			queue = push(queue, node.Left)
		} else {
			output = append(output, node.Val)
			node, queue = pop(queue)
			visited = append(visited, node)

			if node.Right != nil {
				queue = push(queue, node.Right)
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

