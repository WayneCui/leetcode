/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findMode0(root *TreeNode) []int {
	output := []int{}
	if root == nil {
		return output
	}

	memo := make(map[int]int)
	collect(root, memo)

	max := -1
	for n, v := range(memo) {
		if max < v {
			max = v
			output = output[0:0]
			output = append(output, n)
		} else if max == v {
			output = append(output, n)
		}
	}

	return output
}

func collect(root *TreeNode, memo map[int]int) {
	if root == nil {
		return
	}

	memo[root.Val] += 1
	collect(root.Left, memo)
	collect(root.Right, memo)
}

func findMode1(root *TreeNode) []int {
	memo := []int{}
	if root == nil {
		return memo
	}
	
	stack := list.New()
	stack.PushBack(root)

	prevVal := math.MinInt64
	prevCount := 0
	memo = []int{}
	for stack.Len() > 0 {
		node := stack.Back().Value.(*TreeNode)
		fmt.Println(node.Val)

		if node.Left != nil {
			stack.PushBack(node.Left)
			node.Left = nil
		} else {
			if node.Val > prevVal {
				if prevVal != math.MinInt64 {
					if len(memo) == 0 {
						memo = []int{prevCount, prevVal}
					} else if prevCount > memo[0] {
						memo = memo[0:0]
						memo = []int{prevCount, prevVal}
					} else if prevCount == memo[0]{
						memo = append(memo, prevVal)
					} 
				} 
				prevCount = 0
				prevVal = node.Val
				
			} else if node.Val == prevVal {
				prevCount += 1
			}

			node = stack.Remove(stack.Back()).(*TreeNode)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
		}
	}

	if len(memo) == 0 {
		memo = []int{prevCount, prevVal}
	} else if prevCount > memo[0] {
		memo = memo[0:0]
		memo = []int{prevCount, prevVal}
	} else if prevCount == memo[0]{
		memo = append(memo, prevVal)
	} 

	return memo[1:]
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	prevVal := math.MinInt64
	count := 0
	memo := []int{count, prevVal}
	_, _, memo = inorder(root, prevVal, count, memo)
	return memo[1:]
}

func inorder(root *TreeNode, prevVal int, count int, memo []int) (int, int, []int) {
	if root == nil {
		return prevVal, count, memo
	}

	prevVal, count, memo = inorder(root.Left, prevVal, count, memo)

	if prevVal != math.MinInt64 {
		if prevVal == root.Val {
			count += 1
		} else {
			count = 1
		}
	} else {
		count = 1
	}

	if count > memo[0] {
		memo = memo[0:0]
		memo = []int{count, root.Val}
	} else if count == memo[0] {
		memo = append(memo, root.Val)
	}

	prevVal = root.Val

	prevVal, count, memo = inorder(root.Right, prevVal, count, memo)
	return prevVal, count, memo
}

