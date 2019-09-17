/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { return head }

	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		prev = prev.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	if(slow != nil) {
		prev.Next = nil
	}

	left := sortList(head)
	mid := sortList(slow)

	d := merge(left, mid)

	return d.Next
}

func merge(first, second *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for {
		if second == nil {
			curr.Next = first
			break
		} else if first == nil {
			curr.Next = second
			break
		} else if first.Val < second.Val {
			curr.Next = first
			first = first.Next
		} else {
			curr.Next = second
			second = second.Next
		}

		curr = curr.Next
	}

	return dummy
}

