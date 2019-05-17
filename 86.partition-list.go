/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
        return head
    }
    
	low := &ListNode {}
	high := &ListNode {}
	lowHead := low
	highHead := high

	node := head
	for node != nil {
		if node.Val < x {
			low.Next = node
			low = node
		} else {
			high.Next = node
			high = node
		}

		node = node.Next
	}

    high.Next = nil
	low.Next = highHead.Next

	return lowHead.Next
}

