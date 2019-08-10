/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findFrequentTreeSum(root *TreeNode) []int {
	memo := make(map[int]int)
	collect(root, memo)

	res := []int {}
	most := 0
	for key, val := range(memo) {
		if most == val {
			res = append(res, key)
		} else if most < val {
			most = val
			res = []int{key}
		}
	}

	return res
}

func collect(root *TreeNode, memo map[int]int) int {
	if root == nil {
		return 0
	}

	sum := root.Val + collect(root.Left, memo) + collect(root.Right, memo)
	if v, ok := memo[sum]; ok {
		memo[sum] = v + 1
	} else {
		memo[sum] = 1
	}

	return sum
}

