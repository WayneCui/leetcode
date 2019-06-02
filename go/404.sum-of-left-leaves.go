/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
  if root == nil || isLeaf(root){
		return 0
	}

	sum := doSum(root, 0)
	return sum
}

func doSum(root *TreeNode, sum int) int {
	if root != nil {
		if isLeaf(root.Left) {
			sum += root.Left.Val
		} else {
			sum = doSum(root.Left, sum)
		}
	
		if root.Right != nil {
			sum = doSum(root.Right, sum)
		}
	}

	return sum
	
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

