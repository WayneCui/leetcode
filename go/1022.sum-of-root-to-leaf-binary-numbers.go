/*
 * @lc app=leetcode id=1022 lang=golang
 *
 * [1022] Sum of Root To Leaf Binary Numbers
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumRootToLeaf(root *TreeNode) int {
    var buf bytes.Buffer
    col := &[]string{}
    
    _, collection := collect(root, buf, col)
    
    var sum int64
    
    for _, v := range(*collection) {
        iv, _:= strconv.ParseInt(v,2,64)
    
        sum += iv
    }
    
    
    return int(sum)
}

func collect(node *TreeNode, buf bytes.Buffer, col *[]string)  (bytes.Buffer, *[]string) {
    buf.WriteString(strconv.Itoa(node.Val))
    if isLeaf(node)  {
        *col = append(*col, buf.String())
        return buf, col
    } 
    
    if node.Left != nil {
        collect(node.Left, buf, col )
    }
    
    if node.Right != nil {
        collect(node.Right, buf, col)
    }
    
    return buf, col
    
}

func isLeaf(node *TreeNode) bool {
    return node != nil && node.Left == nil && node.Right == nil
}

