/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	n := len(s)
	out := []byte(s)
	partNum := n / k + 1
	for i := 0; i < partNum; i++ {
		if(i % 2 == 0) {
            low := i * k
			high := min((i + 1) * k, n)
            for ; low < high; low, high = low + 1, high + 1 {
				out[low], out[high - 1] = out[high - 1], out[low]
			}
		}
	}

	return string(out)
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
} 
// @lc code=end

