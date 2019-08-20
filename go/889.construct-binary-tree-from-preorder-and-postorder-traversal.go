/*
 * @lc app=leetcode id=889 lang=golang
 *
 * [889] Construct Binary Tree from Preorder and Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main
import (
	"fmt"
)

func constructFromPrePost(pre []int, post []int) *TreeNode {
	n := len(pre)
	if n == 0 { return nil }
	if n == 1 { return &TreeNode { Val: pre[0], }}

	root := &TreeNode {}
	root.Val = pre[0]
	
	sep := indexOf(post, pre[1])
	root.Left = constructFromPrePost(pre[1 : sep + 2], post[0 : sep + 1])
	root.Right = constructFromPrePost(pre[sep + 2 : ], post[sep + 1 : n - 1])

	return root
}

func indexOf(arr []int, target int) int {
	for i, v := range(arr) {
		if target == v {
			return i
		}
	}

	return -1
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	pre := []int{1,2,4,5,3,6,7}
	post := []int{4,5,2,6,7,3,1}
	fmt.Println(constructFromPrePost(pre, post))
}

