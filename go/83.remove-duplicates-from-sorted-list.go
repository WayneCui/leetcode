/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil { return head }
    next := head
    node := head.Next
    for node != nil {
        if node.Val != next.Val {
            next.Next = node
            next = next.Next
        }
        
        node = node.Next
    }
    next.Next = nil
    
    return head
}

