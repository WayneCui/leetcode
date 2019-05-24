/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil { return }
    
    var nodes []*ListNode
    node := head
    for node != nil {
        nodes = append(nodes, node)
        node = node.Next
    }
    
    n := len(nodes)
    for i := 0; i < n / 2; i++ {
        nodes[i].Next = nodes[n - i - 1]
        if n - i - 1 == i + 1 {
            nodes[n - i - 1].Next = nil
        } else {
            nodes[n - i - 1].Next = nodes[i + 1]
            nodes[i + 1].Next = nil
        }
    }
}

