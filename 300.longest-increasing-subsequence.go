/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	memo := make(map[int]int)
	for i := 0; i < n; i++ {
		memo[i] = 0
	}

	memo[0] = 1
	count := 1
	for i := 1; i < n; i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				maxval = max(maxval, memo[j])
			}
		}

		memo[i] = maxval + 1
		count = max(count, memo[i])
	}

	return count
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
