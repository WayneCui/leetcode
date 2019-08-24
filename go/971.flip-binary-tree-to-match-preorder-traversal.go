/*
 * @lc app=leetcode id=971 lang=golang
 *
 * [971] Flip Binary Tree To Match Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var collector []int
	flip(root, voyage, &collector)
	for _, v := range(collector) {
		if v == -1 {
			return []int{ -1}
		}
	}
	return collector
}

func flip(root *TreeNode, voyage []int, collector *[]int) {
	if root.Val != voyage[0] {
		*collector = append(*collector, -1)
	} else {
		if root.Left != nil && root.Right != nil {
			if len(voyage) < 2 {
				*collector = append(*collector, -1)
			}else if root.Left.Val != voyage[1] {
				*collector = append(*collector, root.Val)
				idx := search(voyage[1:], root.Left.Val)
                if idx == -1 {
                    *collector = append(*collector, -1)
                } else {
                    flip(root.Left, voyage[idx + 1:], collector)
				    flip(root.Right, voyage[1:idx + 1], collector)
                }
			} else {
				idx := search(voyage[1:], root.Right.Val)
                if idx == -1 {
                    *collector = append(*collector, -1)
                } else {
                    flip(root.Left, voyage[1:idx + 1], collector)
				    flip(root.Right, voyage[idx + 1:], collector)
                }
			}
		} else if root.Left != nil {
			flip(root.Left, voyage[1:], collector)
		} else if root.Right != nil {
			flip(root.Right, voyage[1:], collector)
		}
	}
}

func search(nums []int, target int) int {
	for i, v := range(nums) {
		if v == target {
			return i
		}
	}

	return -1
}

