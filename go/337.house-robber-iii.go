/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//recursion
func rob0(root *TreeNode) int {
	if root == nil { return 0 }
	a := root.Val
	if root.Left != nil {
		a += rob(root.Left.Left) + rob(root.Left.Right)
	}

	if root.Right != nil {
		a += rob(root.Right.Left) + rob(root.Right.Right)
	}

	b := rob(root.Left) + rob(root.Right)
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//recursion with memo
func rob1(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	return robSub(root, memo)
}

func robSub1(root *TreeNode, memo map[*TreeNode]int) int {
	if root == nil { return 0 }
	if v, ok := memo[root]; ok {
		return v
	}

	a := root.Val
	if root.Left != nil {
		a += robSub(root.Left.Left, memo) + robSub(root.Left.Right, memo)
	}

	if root.Right != nil {
		a += robSub(root.Right.Left, memo) + robSub(root.Right.Right, memo)
	}

	b := robSub(root.Left, memo) + robSub(root.Right, memo)
	val := max(a, b)
	memo[root] = val

	return val
}

//credits to https://leetcode.com/problems/house-robber-iii/discuss/79330/Step-by-step-tackling-of-the-problem
func rob(root *TreeNode) int {
	v1, v2 := robSub(root)
	return max(v1, v2)
}

func robSub(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	l1, l2 := robSub(root.Left)
	r1, r2 := robSub(root.Right)

	return max(l1, l2) + max(r1, r2),  root.Val + l1 + r1
}



