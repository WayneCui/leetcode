/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m >= n {
		return head
	}

	var memo []*ListNode
	memo = append(memo, &ListNode{}) //guard
	memo[0].Next = head

	node := head
	for node != nil {
		memo = append(memo, node)
		node = node.Next
	}

	if m > 1 {
		memo[m - 1].Next = memo[n]
	} else {
		memo[0].Next = memo[n]
	}
	
	if n < len(memo) - 1 {
		memo[m].Next = memo[n + 1]
	} else {
		memo[m].Next = nil
	}
	
	for i := m; i < n; i++ {
		memo[i + 1].Next = memo[i]
	}

	return memo[0].Next
}

