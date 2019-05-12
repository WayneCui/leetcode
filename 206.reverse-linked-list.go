/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    fmt.Println(head)
    if head == nil || head.Next == nil {
        return head
    }
    
	next := head.Next
	curr := head
	var front *ListNode
	for curr != nil  {
		curr.Next = front

		front = curr
		curr = next
        if next != nil {
            next = next.Next
        }        
	}

	return front
}
