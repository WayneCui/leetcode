/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 { return []*TreeNode {} }

	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	trees := buildTrees(nums)

	var output []*TreeNode
	for _, tree := range(trees) {
		root := tree
		output = append(output, &root)
	}

	return output
}

func buildTrees(inorder []int) []TreeNode {
	var output []TreeNode
	n := len(inorder)

	if n == 1 {
		return append(output, TreeNode{ Val: inorder[0], })
	}
	for i, v := range(inorder) {
		var lTmp []TreeNode
		if i > 0 {
			leftSubTrees := buildTrees(inorder[0: i])
			for _, subTree := range(leftSubTrees) {
				left := subTree
				root := TreeNode{
					Val: v,
					Left: &left,
				}

				lTmp = append(lTmp, root)
			}
		} else {
			root := TreeNode{
				Val: v,
			}
			lTmp = append(lTmp, root)
		}
		
		if i < n - 1 {
			rightSubTrees := buildTrees(inorder[i + 1 :])
			for _, subTree := range(rightSubTrees) {
				for _, root := range(lTmp) {
					right := subTree
					root.Right = &right
					output = append(output, root)
				}
			} 
		} else {
			for _, root := range(lTmp) {
				output = append(output, root)
			}
		}
	}

	return output
}

