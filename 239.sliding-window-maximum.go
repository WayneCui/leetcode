/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
func maxSlidingWindow1(nums []int, k int) []int {
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

// another way, credits to https://leetcode.com/problems/sliding-window-maximum/discuss/65881/O(n)-solution-in-Java-with-two-simple-pass-in-the-array
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	output := []int{}
	if n == 0 { return output }

	forward := make([]int, n)
	backward := make([]int, n)

	epoch := 1
	currMax := 0
	for i := 0; i < n; i++ {
		if i < epoch * k {
			currMax = max(currMax, nums[i])
		} else {
			currMax = nums[i]
			epoch += 1
		}
		forward[i] = currMax
	}

	currMax = 0
	epoch -= 1
	for i := n - 1; i >= 0; i-- {
		if i > epoch * k {
			currMax = max(currMax, nums[i])
		} else {
			currMax = nums[i]
			epoch -= 1
		}

		backward[i] = currMax
	}

	for i := 0; i < n - k + 1; i++ {
		output = append(output, max(forward[i + k - 1], backward[i]))
	}

	return output
}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}
