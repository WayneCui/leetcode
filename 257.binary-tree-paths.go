/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var collector []string

	if root != nil {
		var buffer bytes.Buffer
		buffer.WriteString(strconv.Itoa(root.Val))
		reduce(root.Left, buffer, &collector)
		reduce(root.Right, buffer, &collector)

		if root.Left == nil && root.Right == nil {
			collector = append(collector, buffer.String())
		}
	}

	return collector
}

func reduce(root *TreeNode, buffer bytes.Buffer, collector *[]string) {
	if root == nil {
		return
	}

	buffer.WriteString("->")
	buffer.WriteString(strconv.Itoa(root.Val))
	if root.Left != nil {
		reduce(root.Left, buffer, collector)
	}

	if root.Right != nil {
		reduce(root.Right, buffer, collector)
	}

	if root.Left == nil && root.Right == nil {
		*collector = append(*collector, buffer.String())
	}
}

