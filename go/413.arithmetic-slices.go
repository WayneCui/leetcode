/*
 * @lc app=leetcode id=413 lang=golang
 *
 * [413] Arithmetic Slices
 */
func numberOfArithmeticSlices0(A []int) int {
	n := len(A)
	if n < 3 { return 0 }

	dp := make([]int, n)
	start := 0
	count := 0
	for i := 2; i < n; i++ {
		if A[i] - A[i - 1] == A[i - 1] - A[i - 2] {
			dp[i] = i - 2 - start + 1
			count += dp[i]
		} else {
			start = i - 1
		}
	}

	return count
}

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n < 3 { return 0 }

	start := 0
	count := 0
	for i := 2; i < n; i++ {
		if A[i] - A[i - 1] == A[i - 1] - A[i - 2] {
			count += (i - 2 - start + 1)
		} else {
			start = i - 1
		}
	}

	return count
}
