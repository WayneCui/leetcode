/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */
 func rob(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }
	if n == 1 { return nums[0]}
	return max(toRob(nums[0:n - 1]),toRob(nums[1:n]))
}

func toRob(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }
	if n == 1 { return nums[0]}

	memo := make([]int, n)
	memo[0] = nums[0]
	memo[1] = max(memo[0], nums[1])
	for i := 2; i < n; i++ {
		memo[i] = max(memo[i - 2] + nums[i], memo[i - 1])
	}

	return memo[n - 1]
}

func max(a, b int ) int {
	if a > b {
		return a
	} else {
		return b
	}
}
