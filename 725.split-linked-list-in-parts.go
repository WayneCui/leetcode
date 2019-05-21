/*
 * @lc app=leetcode id=725 lang=golang
 *
 * [725] Split Linked List in Parts
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func splitListToParts(root *ListNode, k int) []*ListNode {
    var result []*ListNode
    for i := 0; i < k; i++ {
        result = append(result, nil)
    }
    
    var nodes []*ListNode
    node := root
    for node != nil {
		nodes = append(nodes, node)
		node = node.Next
    }
    
    n := len(nodes)
    perGrpup := n / k
	extra := n % k

	idxNodes := 0
	idxResult := 0
	for idxNodes < n {
		result[idxResult] = nodes[idxNodes]
		if extra > 0 {
			idxNodes = idxNodes + perGrpup + 1
			extra -= 1
		} else {
			idxNodes = idxNodes + perGrpup
		}
		nodes[idxNodes - 1].Next = nil
		idxResult += 1
	}
	
    return result
}

