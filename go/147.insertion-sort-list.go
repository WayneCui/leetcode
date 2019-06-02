/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	superHead := &ListNode {
		Val: math.MinInt32,
		Next: nil,
	}

	node := head
	for node != nil {
		next := node.Next
		pos := findPos(superHead, nil, node)

		tmp := pos.Next
		pos.Next = node
		node.Next = tmp

		node = next
	}

	return superHead.Next
}

func findPos(from *ListNode, to *ListNode, node *ListNode) *ListNode {
	for from.Next != to {
		fast := from
		slow := from
		for fast != to && fast.Next != to {
			slow = slow.Next
			fast = fast.Next.Next
		}

		if node.Val > slow.Val {
			from = slow
		} else if node.Val < slow.Val {
			to = slow
		} else {
			return slow
		}
	}
	return from
}

