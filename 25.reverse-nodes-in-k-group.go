/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	superHead := ListNode {
		Next: head,
	}

	curr := &superHead
	front := head
	end := head
	for i := 0; i < k - 1; i++ {
		if end != nil {
			end = end.Next
		} else {
			break
		}
	}
		
	for end != nil {
		curr.Next = end
		theEnd := end.Next
		reverse(front, theEnd)
		front.Next = theEnd

		curr = front
		front = front.Next
		end = front
		for i := 0; i < k - 1; i++ {
			if end != nil {
				end = end.Next
			} else {
				end = nil
				break
			}
		}
	}

	return superHead.Next
}

func reverse(front *ListNode, end *ListNode) {
	prev := front
	behind := front.Next
	for behind != end {
		tmp := behind.Next
		behind.Next = prev
		prev = behind
		behind = tmp
	}
}

