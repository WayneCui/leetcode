/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor0(root, p, q *TreeNode) *TreeNode {
	_, ancestor := lca(root, p, q)
	return ancestor
}

func lca(root, p, q *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	var count int
	var ancestor *TreeNode

	lc, ln := lca(root.Left, p, q)
	rc, rn := lca(root.Right, p, q)
	if lc == 2 {
		ancestor = ln
	} else if rc == 2 {
		ancestor = rn
	}
    
    if root == p {
		count += 1
	} else if root == q {
		count += 1
	}
	
	count += lc + rc

	if count == 2 && ancestor == nil {
		ancestor = root
	} 

	return count, ancestor
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var pPath []*TreeNode
	var qPath []*TreeNode
	findPath(root, p, &pPath)
	findPath(root, q, &qPath)

	np := len(pPath)
	nq := len(qPath)
	ancestor := root
	for np > 0 && nq > 0 {
		if pPath[np - 1] == qPath[nq - 1] {
			ancestor = pPath[np - 1]
		}

		np -= 1
		nq -= 1
	}

	return ancestor
}

func findPath(root, target *TreeNode, memo *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	existLeft := findPath(root.Left, target, memo)
	existRight := findPath(root.Right, target, memo)
	if target == root || existLeft || existRight {
		*memo = append(*memo, root)
		return true
	}

	return false
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}

	return nil
}

