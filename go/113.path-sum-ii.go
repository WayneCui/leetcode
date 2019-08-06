/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// stuck by
//  [-6,null,-3,-6,0,-6,-5,4,null,null,null,1,7]
//  -21
 func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil{
		return [][]int{}
	}
    container := [][]int{}
	path := []int {}
	collect(root, path, sum, &container)

	return container
}

func collect(root *TreeNode, path []int, sum int, container *[][]int) {
	if root.Left == nil && root.Right == nil {        
		if root.Val == sum {
            path = append(path, root.Val)
            res := make([]int, len(path))
            copy(res, path)
            *container = append(*container, res)
		}
		return
	}

	path = append(path, root.Val)
    sum -= root.Val
    
    if root.Left != nil {
		collect(root.Left, path, sum, container)
	}
    
	if root.Right != nil {
		collect(root.Right, path, sum, container)
	}
}


