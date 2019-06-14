/*
 * @lc app=leetcode id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str0(t *TreeNode) string {
	if t == nil {
		return ""
	}

	buf := new(strings.Builder)
	buf.WriteString(strconv.Itoa(t.Val))
	if t.Left != nil {
		buf.WriteString("(")
		buf.WriteString(tree2str(t.Left))
		buf.WriteString(")")
	} else {
		if t.Right != nil {
			buf.WriteString("()")
		}
	}
	
	if t.Right != nil {
		buf.WriteString("(")
		buf.WriteString(tree2str(t.Right))
		buf.WriteString(")")
	}
	
	return buf.String()
}

//iterative
func tree2str(t *TreeNode) string {
	
	if t == nil {
		return ""
	}

	buf := new(strings.Builder)
	stack := []*TreeNode { t, t }
	var node *TreeNode
	for len(stack) > 0 {
		node, stack = pop(stack)
		if len(stack) > 0 && peek(stack) == node {
			buf.WriteString("(")
			buf.WriteString(strconv.Itoa(node.Val))
			if node.Left == nil && node.Right != nil {
				buf.WriteString("()")
			}

			if node.Right != nil {
				stack = push(stack, node.Right)
				stack = push(stack, node.Right)
			}
			
			if node.Left != nil {
				stack = push(stack, node.Left)
				stack = push(stack, node.Left)
			}	
		} else {
			buf.WriteString(")")
		}
	} 

	if buf.Len() == 0 { return "" }

	return buf.String()[1 : buf.Len() - 1]
}

func pop(a []*TreeNode) (poped *TreeNode, remained []*TreeNode) {
	x, a := a[len(a)-1], a[:len(a)-1]
	return x, a
}

func peek(a []*TreeNode) *TreeNode {
	return a[len(a)-1]
}

func push(a []*TreeNode, node *TreeNode) []*TreeNode {
	return append(a, node)
}

func tree2str2(t *TreeNode) string {
	
	if t == nil {
		return ""
	}

	buf := new(strings.Builder)

	stack := list.New()
	stack.PushBack(t)
	var node *TreeNode
	for stack.Len() > 0 {
		node = stack.Remove(stack.Back()).(*TreeNode)
		if stack.Len() > 0 && stack.Back().Value.(*TreeNode) == node {
			buf.WriteString(strconv.Itoa(node.Val))
			buf.WriteString("(")
			if node.Left == nil && node.Right != nil {
				buf.WriteString("()")
			}

			if node.Right != nil {
				stack.PushBack(node.Right)
				stack.PushBack(node.Right)
			}
			
			if node.Left != nil {
				stack.PushBack(node.Left)
				stack.PushBack(node.Left)
			}	
		} else {
			buf.WriteString(")")
		}
	} 

	return buf.String()
}

