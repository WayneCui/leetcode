/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */
func findKthLargest(nums []int, k int) int {
	h := IntHeap {}
	for _, num := range nums {
		heap.Push(&h, num)
	}
	fmt.Println(h)

	var res int
	for i := 0; i < k; i++ {
		res = heap.Pop(&h).(int)
	}

	return res
}

type IntHeap []int
func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
} 

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[ n - 1]
	*h = old[0 : n - 1]
	return x
}

