/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 */
func lastStoneWeight(stones []int) int {
	h := IntHeap([]int{})
	for _, val := range(stones) {
		heap.Push(&h, val)
	}

	for len(h) > 1 {
		l := heap.Pop(&h).(int)
		s := heap.Pop(&h).(int)
		diff := l - s
		if diff > 0 {
			heap.Push(&h, diff)
		}
	}

	if len(h) == 0 { 
		return 0 
	} else {
		return h[0]
	}
}

type IntHeap []int 
func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}
