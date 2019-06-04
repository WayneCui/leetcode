/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST0(root *TreeNode) bool {
		inorderView := inorderTraversal(root)
		n := len(inorderView)
		for i := 0; i < n - 1; i++ {
			if inorderView[i] >= inorderView[i + 1] {
				return false
			}
		}

		return true
}

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


func isValidBST(root *TreeNode) bool {
	uppper := math.MaxInt64
	lower := math.MinInt64
	return helper(root, uppper, lower)
}

func helper(root *TreeNode, max int, min int) bool {
	if root == nil { return true }

	if root.Val >= max || root.Val <= min {
		return false
	} 

	return helper(root.Left, root.Val, min) && helper(root.Right, max, root.Val)
}
 
