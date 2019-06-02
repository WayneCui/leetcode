/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	superHead := &ListNode {}
	if len(lists) == 0 {
		return nil
	}
	theHeap := NodeHeap{}
	for _, node := range(lists) {
		if node != nil {
			heap.Push(&theHeap, node)
		}
	}
	heap.Init(&theHeap)

	curr := superHead
	for len(theHeap) > 0 {
		node := heap.Pop(&theHeap).(*ListNode)
		if node.Next != nil {
			heap.Push(&theHeap, node.Next)
		}

		curr.Next = node
		curr = curr.Next
	}

	return superHead.Next
}

type NodeHeap []*ListNode
func (theHeap NodeHeap) Len() int { return len(theHeap) }
func (theHeap NodeHeap) Less(i, j int) bool {
	return theHeap[i].Val < theHeap[j].Val
}
func (theHeap NodeHeap) Swap(i, j int) { theHeap[i], theHeap[j] = theHeap[j], theHeap[i] }
func (theHeap *NodeHeap) Push(x interface{}) {
    *theHeap = append(*theHeap, x.(*ListNode))
}

func (theHeap *NodeHeap) Pop() interface{} {
	old := *theHeap
	n := len(old)
	x := old[n - 1]
	*theHeap = old[0: n - 1]
	return x
}

