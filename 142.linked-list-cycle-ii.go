/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast := head
	slow := head

	for fast != nil &&  fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

