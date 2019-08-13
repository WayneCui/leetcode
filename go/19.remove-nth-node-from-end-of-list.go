/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode {
		Next: head,
	}
	first := dummy
	second := dummy
	for i := 1; i <= n + 1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}

