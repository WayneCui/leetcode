/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumNumbers(root *TreeNode) int {
    if root == nil { return 0 }
	collector := [][]int{}
	init := []int{}
	traversal(root, init, &collector)

	sum := 0
	for _, p := range(collector) {
		sum += combine(p)
	}

	return sum
}

func traversal(root *TreeNode, path []int, collector *[][]int) {
	if root.Left == nil && root.Right == nil {
		p := make([]int, len(path))
		copy(p, path)
		p = append(p, root.Val)
		*collector = append(*collector, p)
		return 
	}

	path = append(path, root.Val)
	if(root.Left != nil) {
		traversal(root.Left, path, collector)
	}

	if(root.Right != nil) {
		traversal(root.Right, path, collector)
	}
}

func combine(p []int) int {
	var buffer bytes.Buffer
	for _, v := range(p) {
		buffer.WriteString(strconv.Itoa(v))
	}

	if v, err := strconv.Atoi(buffer.String()); err == nil {
		return v
	}

	return 0
}


//a more memory efficient way
 func sumNumbers(root *TreeNode) int {
    if root == nil { return 0 }
	sum := []int{0}
	init := []int{}
	traversal(root, init, &sum)

	return sum[0]
}

func traversal(root *TreeNode, path []int, sum *[]int) {
	if root.Left == nil && root.Right == nil {
		p := make([]int, len(path))
		copy(p, path)
		p = append(p, root.Val)
        (*sum)[0] += combine(p)
		return 
	}

	path = append(path, root.Val)
	if(root.Left != nil) {
		traversal(root.Left, path, sum)
	}

	if(root.Right != nil) {
		traversal(root.Right, path, sum)
	}
}


