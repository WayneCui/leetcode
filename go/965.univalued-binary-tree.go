/*
 * @lc app=leetcode id=965 lang=golang
 *
 * [965] Univalued Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isUnivalTree(root *TreeNode) bool{
    
    return validate(root, true)
}

func validate(node *TreeNode, state bool) bool {
    if node == nil {
        return state
    }
    
    if node.Left != nil {
        state = state && (node.Val == node.Left.Val)
    }
    
    if node.Right != nil {
        state = state && (node.Val == node.Right.Val)
    }
    
    return state && validate(node.Left, state) && validate(node.Right, state)
}

