/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	superHead := &ListNode{
		Next: head.Next,
	}
	front := head.Next
	back := head
	curr := superHead
	for front != nil {
		curr.Next = front
		back.Next = front.Next
		front.Next = back

		curr = back
		back = back.Next
		if back != nil {
			front = back.Next
		} else {
			front = nil
		}
	}

	return superHead.Next
}

