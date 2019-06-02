/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    
    super := &ListNode {
        Next : head,
    }
    odd := head
    even := head.Next
    second := head.Next
    
    for even != nil && even.Next != nil{
        odd.Next = even.Next    
        if odd.Next != nil {
            odd = odd.Next
		}
		
        even.Next = odd.Next
        if even.Next != nil {
            even = even.Next
        }
        
    }

    odd.Next = second
    
    return super.Next
    
}

