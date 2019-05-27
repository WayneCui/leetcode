/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	output := []int{}
	if n == 0 { return output }

	for i := 0; i < n - k + 1; i++ {
		h := IntHeap {}
		for j := 0; j < k; j++ {
			heap.Push(&h, nums[i + j])
		}
		output = append(output, h[0])
	}

	return output
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

