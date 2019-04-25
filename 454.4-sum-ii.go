/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 */
func fourSumCount(A []int, B []int, C []int, D []int) int {
	n := len(A)
	memo := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v, ok := memo[A[i] + B[j]]
			if ok {
				memo[A[i] + B[j]] = v + 1
			} else {
				memo[A[i] + B[j]] = 1
			}
		}
	}

	output := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v, ok := memo[-C[i] - D[j]]
			if ok {
				output += v
			}
		}
	}

	return output

}

