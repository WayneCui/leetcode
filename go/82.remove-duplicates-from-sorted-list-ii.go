/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	superHead := &ListNode {
		Next: head,
	}
	
	target := head
	next := superHead
	node := head.Next
	count := 0
	for node != nil {
		if node.Val == target.Val {
			count += 1
			node = node.Next
			continue
		}

		if count == 0 {
			next.Next = target
			next = next.Next
		}
		
		target = node
		node = node.Next
		count = 0
	}

	if target.Next == nil {
		next.Next = target
	} else {
		next.Next = nil
	}

	return superHead.Next
}

