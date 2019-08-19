/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil { return l2 }
	if l2 == nil { return l1 }

	head := &ListNode {}
	if l1.Val <= l2.Val {
		head.Val = l1.Val
		head.Next = mergeTwoLists(l1.Next, l2)
	} else {
		head.Val = l2.Val
		head.Next = mergeTwoLists(l1, l2.Next)
	}

	return head
}

