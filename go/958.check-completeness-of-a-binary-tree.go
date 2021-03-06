/*
 * @lc app=leetcode id=958 lang=golang
 *
 * [958] Check Completeness of a Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    layer := []*TreeNode { root }

    end := false
    for len(layer) > 0 {
        nextLayer := []*TreeNode {}
        for _, node := range(layer) {
            if node == nil {
                end = true
            } else {
                if end {
                    return false
                }
                nextLayer = append(nextLayer, node.Left)
                nextLayer = append(nextLayer, node.Right)
            }
        }

        layer = nextLayer
    }

    return true
}

// a more clear way
func isCompleteTree(root *TreeNode) bool {
    layer := []*TreeNode { root }
    prev := &TreeNode {}
    isComplete := true
    for len(layer) > 0 {
        nextLayer := []*TreeNode {}
        for _, node := range(layer){
            isComplete = isComplete && !(prev == nil && node != nil)

            if node != nil {
                nextLayer = append(nextLayer, node.Left)
                nextLayer = append(nextLayer, node.Right)
            }
            
            prev = node
        }
        
        layer = nextLayer
    }
    
    return isComplete
}

