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
			node.Left = nil
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
//morris traversal
func inorderTraversal2(root *TreeNode) []int {
	output := []int{}
	curr := root
	var prev *TreeNode

	for curr != nil {
		if curr.Left != nil {
			prev = curr.Left

			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}

			if prev.Right == nil {
				prev.Right = curr
				curr = curr.Left
			} else { //prev.Right == curr
				output = append(output, curr.Val)

				prev.Right = nil
				curr = curr.Right
			}
		} else {
			output = append(output, curr.Val)
			curr = curr.Right
		}
	}

	return output
}

//another iteratively way, will destory the original tree
func inorderTraversal3(root *TreeNode) []int {
	var output []int
	if root == nil {
		return output
	}
	
	stack := []*TreeNode{ root }
	for len(stack) > 0 {
		node := peek(stack)

		if node.Left != nil {
			stack = push(stack, node.Left)
			node.Left = nil
		} else {
			output = append(output, node.Val)
			node, stack = pop(stack)

			if node.Right != nil {
				stack = push(stack, node.Right)
			}
		}
	}

	return output
}

//again, another iteratively way, will destory the original tree
func inorderTraversal(root *TreeNode) []int {
	var output []int
	if root == nil {
		return output
	}
	
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Back().Value.(*TreeNode)

		if node.Left != nil {
			stack.PushBack(node.Left)
			node.Left = nil
		} else {
			output = append(output, node.Val)
			node = stack.Remove(stack.Back()).(*TreeNode)

			if node.Right != nil {
				stack.PushBack(node.Right)
			}
		}
	}

	return output
}

