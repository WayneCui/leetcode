/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST_1(head *ListNode) *TreeNode {
	if head == nil { return nil }
	return toBst(head, nil)
}

func toBst(head *ListNode, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}

	slow := head
	fast := head

	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}

	thead := &TreeNode {
		Val: slow.Val,
		Left: toBst(head, slow),
		Right: toBst(slow.Next, tail),
	}

	return thead
}

func sortedListToBST(head *ListNode) *TreeNode {
	var nums []int
	for {
		if head == nil {
			break
		}
		nums = append(nums, head.Val)
		head = head.Next
	}

	return sortedArrayToBST(nums)
	
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode { Val: nums[0],}
	} else if n == 2 {
		return &TreeNode {
			Val: nums[0],
			Right: &TreeNode {
				Val: nums[1],
			},
		}
	}

	mid := n / 2
	root := TreeNode{
		Val: nums[mid],
		Left: sortedArrayToBST(nums[0: mid]),
		Right: sortedArrayToBST(nums[mid + 1:]),
	}

	return &root
}


