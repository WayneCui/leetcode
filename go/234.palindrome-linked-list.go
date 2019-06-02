/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return true
	}

	mid := middleNode(head)
	revRight := reverseList(mid)
	node := head
	for node != mid {
		if node.Val != revRight.Val {
			return false
		}

		node = node.Next
		revRight = revRight.Next
	}

	return true
}

func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    return slow
}

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
	next := head.Next
	curr := head
	var front *ListNode
	for curr != nil  {
		curr.Next = front

		front = curr
		curr = next
        if next != nil {
            next = next.Next
        }        
	}

	return front
}

