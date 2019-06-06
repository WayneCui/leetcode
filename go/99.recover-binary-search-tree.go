/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// inorder traversal
func recoverTree0(root *TreeNode)  {
	nodesInorder := inorderTraversal(root)
	
	n := len(nodesInorder)
	var target1, target2 *TreeNode
	for i := 0; i < n - 1; i++ {
		if nodesInorder[i].Val > nodesInorder[i + 1].Val {
			if target1 == nil {
				target1 = nodesInorder[i]
			}
			
			target2 = nodesInorder[i + 1]
		}
	}

	if target1 != nil && target2 != nil {
		target1.Val, target2.Val = target2.Val, target1.Val
	}
}

func inorderTraversal(root *TreeNode) []*TreeNode {
    if root == nil { 
		return []*TreeNode{}
	}

	output := []*TreeNode{}
	if root.Left != nil {
		output = append(output, inorderTraversal(root.Left)...)
	}

	output = append(output, root)

	if root.Right != nil {
		output = append(output, inorderTraversal(root.Right)...)
	}

	return output
}


