/*
 * @lc app=leetcode id=1019 lang=golang
 *
 * [1019] Next Greater Node In Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
type Stack struct {
    list *list.List
}

func NewStack() *Stack {
    list := list.New()
    return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
    stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
    e := stack.list.Back()
    if e != nil {
        stack.list.Remove(e)
        return e.Value
    }
    return nil
}

func (stack *Stack) Peak() interface{} {
    e := stack.list.Back()
    if e != nil {
        return e.Value
    }

    return nil
}

func (stack *Stack) Len() int {
    return stack.list.Len()
}

func (stack *Stack) Empty() bool {
    return stack.list.Len() == 0
}

func nextLargerNodes0(head *ListNode) []int {
    if head == nil {
        return []int{}
    } 
    
    if head.Next == nil {
        return []int{0}
    }
    output := []int{}
    stack := NewStack()
    
    node := head
    for node != nil {
        c := 0
        for {
            prev := stack.Peak()
            if prev == nil { break }
            if prev.(int) < node.Val {
                stack.Pop()
                c += 1
            } else {
                break
            }
        }
        
        n := len(output)
        for c > 0 {
            if output[n - 1] == 0 {
                output[n - 1] = node.Val
                c -= 1
                
            }
            n -= 1
        }
        
        
        stack.Push(node.Val)
        output = append(output, 0)
        node = node.Next
    }
    
    return output
}


func nextLargerNodes(head *ListNode) []int {
    if head == nil {
        return []int{}
    } 
    
    if head.Next == nil {
        return []int{0}
    }

    var nodeVals []int
    node := head
    for node != nil {
        nodeVals = append(nodeVals, node.Val)
        node = node.Next
    }
    n := len(nodeVals)

    output := make([]int, n)
    stack := NewStack()
    
    for i := 0; i < n; i++ {
        for {
            if !stack.Empty() && nodeVals[stack.Peak().(int)] < nodeVals[i] {
                output[stack.Pop().(int)] = nodeVals[i]
            } else {
                break
            }
        }

        stack.Push(i)
    }    
    
    return output
}

