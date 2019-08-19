/*
 * @lc app=leetcode id=865 lang=golang
 *
 * [865] Smallest Subtree with all the Deepest Nodes
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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, node := helper(root)
	return node
}

func helper(root *TreeNode) (depth int, node *TreeNode) {
	if root == nil { return 0, root }
	
	dl, nl := helper(root.Left)
	dr, nr := helper(root.Right)
	if dl == dr {
		return dl + 1, root
	} else if dl > dr {
		return dl + 1, nl
	} else {
		return dr + 1, nr
	}
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode {
		Val: 1,
		Left: &TreeNode {
			Val: 2,
			Left: &TreeNode {
				Val: 4,
			},
			Right: &TreeNode {
				Val: 5,
			},
		},
		Right: &TreeNode {
			Val: 3,
		},
	}

	fmt.Println(subtreeWithAllDeepest(root))
}

