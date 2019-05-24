/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */
type KthLargest struct {
	K int
    TheHeap IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest {
		K: k,
		TheHeap: IntHeap([]int{}),
	}
	for _, v := range(nums) {
		if len(kl.TheHeap) < k {
			heap.Push(&kl.TheHeap, v)
		} else {
			theSmall := kl.TheHeap[0]
			if theSmall < v {
				heap.Pop(&kl.TheHeap)
				heap.Push(&kl.TheHeap, v)
			} 
		}
	}
	return kl
}


func (this *KthLargest) Add(val int) int {
	if len(this.TheHeap) < this.K {
		heap.Push(&this.TheHeap, val)
	} else {
		theSmall := this.TheHeap[0]
		if theSmall < val {
			heap.Pop(&this.TheHeap)
			heap.Push(&this.TheHeap, val)
		}
	}

	return this.TheHeap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

 type IntHeap []int
 func (h IntHeap) Len() int { return len(h)}
 func (h IntHeap) Less(i, j int) bool { return h[i] < h[j]}
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

