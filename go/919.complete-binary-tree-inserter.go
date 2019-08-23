/*
 * @lc app=leetcode id=919 lang=golang
 *
 * [919] Complete Binary Tree Inserter
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	Nodes []*TreeNode
}


func Constructor(root *TreeNode) CBTInserter {
	inserter := new(CBTInserter)
	
	nodes := []*TreeNode { root }
	from := 0
	to := 1
	for from < to {
		for i := from; i < to; i++ {
			node := nodes[i]
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			} else {
				break
			}
	
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			} else {
				break
			}
		}

		from = to
		to = len(nodes)
	}

	inserter.Nodes = nodes
	return *inserter
}


func (this *CBTInserter) Insert(v int) int {
	n := len(this.Nodes)

	c := &TreeNode {
		Val: v,
	}
	this.Nodes = append(this.Nodes, c)

	var p *TreeNode
	if n % 2 == 0 {
		p = this.Nodes[(n - 2) / 2]
		p.Right = c
		
	} else {
		p = this.Nodes[(n - 1) / 2]
		p.Left = c
	}
	return p.Val
}


func (this *CBTInserter) Get_root() *TreeNode {
    return this.Nodes[0]
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */

