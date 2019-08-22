/*
 * @lc app=leetcode id=894 lang=golang
 *
 * [894] All Possible Full Binary Trees
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT0(N int) []*TreeNode {
	res := []*TreeNode {}
    
	if N % 2 == 0 {
		return res
	}

	if N == 1 {
		res = append(res, &TreeNode{ Val: 0, })
		return res
	}

	for i := 1; i <= N - 2; i += 2 {
		for _, lnode := range(allPossibleFBT(i)) {
			for _, rnode := range(allPossibleFBT( N - 1 - i)){
				root := &TreeNode {
					Val: 0,
					Left: lnode,
					Right: rnode,
				}

				res = append(res, root)
			}
		}
	}	
	return res
}


//with cache
func allPossibleFBT(N int) []*TreeNode {
	if N % 2 == 0 {
		return []*TreeNode {}
	}
	cache := make(map[int][]*TreeNode)
	return findPossible(N, cache)
}

func findPossible(N int, cache map[int][]*TreeNode) []*TreeNode {
	if v, ok := cache[N]; ok {
		return v
	}

	res := []*TreeNode {}
	if N == 1 {
		res = append(res, &TreeNode{ Val: 0, })
		return res
	}

	for i := 1; i <= N - 2; i += 2 {
		for _, lnode := range(findPossible(i, cache)) {
			for _, rnode := range(findPossible( N - 1 - i, cache)){
				root := &TreeNode {
					Val: 0,
					Left: lnode,
					Right: rnode,
				}

				res = append(res, root)
			}
		}
	}	

	cache[N] = res
	return res
}

