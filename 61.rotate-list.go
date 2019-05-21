/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1
	node := head
	for node.Next != nil {
		n += 1
		node = node.Next
	}

	end := node
	end.Next = head

	r := k % n
	nd := head
	for i := 0; i < n - r - 1; i++ {
		nd = nd.Next
	}

	result := nd.Next
	nd.Next = nil

	return result
}

