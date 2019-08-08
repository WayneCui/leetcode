/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type BSTIterator struct {
	vals []int
	idx int
}


func Constructor(root *TreeNode) BSTIterator {
	stack := list.New()
	vals := []int{}
	
	if root != nil { 
		stack.PushBack(root)
		for stack.Len() > 0 {
			curr := stack.Back().Value.(*TreeNode)
			if curr.Left != nil {
				stack.PushBack(curr.Left)
				curr.Left = nil
			} else {
				vals = append(vals, curr.Val)
	
				curr := stack.Remove(stack.Back()).(*TreeNode)
				if curr.Right != nil {
					stack.PushBack(curr.Right)
				}
			}
		}
	}
    
    iter := new(BSTIterator)
	iter.vals = vals
	iter.idx = 0

	
    return *iter
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    v := this.vals[this.idx]
    this.idx += 1
	return v
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return this.idx < len(this.vals)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

// a more efficient version
type BSTIterator struct {
	stack *list.List
}


func Constructor(root *TreeNode) BSTIterator {
	stack := list.New()
	
	if root != nil { 
		stack.PushBack(root)
		for {
			curr := stack.Back().Value.(*TreeNode)
			if curr.Left != nil {
				stack.PushBack(curr.Left)
			} else {
				break
			}
		}
	}
    
    iter := BSTIterator{
		stack: stack,
	}	
    return iter
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack.Remove(this.stack.Back()).(*TreeNode)
	if node.Right != nil {
		this.stack.PushBack(node.Right)
		for {
			curr := this.stack.Back().Value.(*TreeNode)
			if curr.Left != nil {
				this.stack.PushBack(curr.Left)
			} else {
				break
			}
		}
	}

	return node.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return this.stack.Len() > 0 
}

