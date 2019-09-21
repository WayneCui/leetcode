/*
 * @lc app=leetcode id=1145 lang=golang
 *
 * [1145] Binary Tree Coloring Game
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Credits to https://leetcode.com/problems/binary-tree-coloring-game/discuss/350570/JavaC%2B%2BPython-Simple-recursion-and-Follow-Up
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var left, right int
	count(root, x, &left, &right)
	return max(left, right, n - left - right - 1) > n / 2
}

func count(root *TreeNode, x int, left *int, right *int) int {
	if root == nil { return 0}
	l := count(root.Left, x, left, right)
	r := count(root.Right, x, left, right)
	if root.Val == x {
		*left = l
		*right = r
	}

	return l + r + 1
}

func max(a...int) int {
	m := math.MinInt32
	for _, x := range a {
		if x > m {
			m = x
		}
	}

	return m
}

