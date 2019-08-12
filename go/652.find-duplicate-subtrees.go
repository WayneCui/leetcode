/*
 * @lc app=leetcode id=652 lang=golang
 *
 * [652] Find Duplicate Subtrees
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil { return []*TreeNode{} }
	memo := make(map[string]int)
	res := []*TreeNode {}
	traversal(root, memo, &res)

	return res
}

func traversal(root *TreeNode, memo map[string]int, res *[]*TreeNode) string  {
	if root == nil { return "#" }

	seri := strconv.Itoa(root.Val) + "," + traversal(root.Left, memo, res) + "," + traversal(root.Right, memo, res)
	if v, ok := memo[seri]; ok {
		if v == 1 { 
			*res = append(*res, root)
		}
		
		memo[seri] = memo[seri] + 1
	} else {
		memo[seri] = 1
	}
	return seri
}

