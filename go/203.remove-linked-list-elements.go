/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode {
		Next: head,
	}
	front := newHead
	curr := head
	for curr != nil {
		if curr.Val == val {
			front.Next = curr.Next
			curr = curr.Next
		} else {
			front = curr
			curr = curr.Next
		}
	}

	return newHead.Next
}

