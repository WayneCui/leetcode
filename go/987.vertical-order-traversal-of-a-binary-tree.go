/*
 * @lc app=leetcode id=987 lang=golang
 *
 * [987] Vertical Order Traversal of a Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func verticalTraversal(root *TreeNode) [][]int {
	memo := make(map[int]map[int][]int)
	dsf(root, 0, 0, memo)

	var keys []int
	for key, _ := range memo {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	var output [][]int
	for _, k := range(keys) {
		var vals []int

		var keys2 []int
		for key, _ := range memo[k] {
			keys2 = append(keys2, key)
		}
        
        sort.Slice(keys2, func(i, j int) bool {
            return keys2[i] > keys2[j]
        })
		for _, k2 := range keys2 {
			vals = append(vals, memo[k][k2]...)
		}

		output = append(output, vals)
	}

	return output
}

func dsf(root *TreeNode, horizontal int, vertical int, memo map[int]map[int][]int) {
	if(root == nil) { return }
	if m, ok := memo[horizontal]; ok {
		if s, ok := m[vertical]; ok {
			memo[horizontal][vertical] = insertByAsc(s, root.Val)
		} else {
			memo[horizontal][vertical] = []int{root.Val}
		}
	} else {
		memo[horizontal] = make(map[int][]int)
		memo[horizontal][vertical] = []int{root.Val}
	}

	dsf(root.Left, horizontal - 1, vertical - 1, memo)
	dsf(root.Right, horizontal + 1, vertical - 1, memo)
}

func insertByAsc(s []int, new int) []int {
    fmt.Println(s)
    low := 0
    high := len(s)
    for low < high {
        mid := (high - low) / 2 + low
        if (s[mid] < new) {
            low = mid + 1
        } else {
            high = mid
        }
    }
    
    return append(s[0:low], append([]int{new}, s[low:]...)...)
}

