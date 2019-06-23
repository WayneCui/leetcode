/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob1(nums []int) int {
	memo := make([]int, len(nums))
	for i := range(nums) {
		memo[i] = -1
	}
	return toRob(nums, len(nums) - 1, &memo)
}

func toRob(nums []int, i int, memo *[]int) int {
	if i < 0 {
		return 0
	}

    if (*memo)[i] != -1 {
        return (*memo)[i]
	}

	result := max(toRob(nums, i - 2, memo) + nums[i], toRob(nums, i - 1, memo))
    (*memo)[i] = result
	return result
}


func rob2(nums []int) int {
	n := len(nums)
	if n == 0 { return 0}
	memo := make([]int, n + 1)
	memo[0] = 0
	memo[1] = memo[0]
	for i := 1; i < n; i++ {
		memo[i + 1] = max(memo[i], memo[i - 1] + nums[i])
	}

	return memo[n]
}


func rob(nums []int) int {
	n := len(nums)
	if n == 0 { return 0}
	prev1 := 0
	prev2 := 0
	for _, num := range(nums) {
		prev1, prev2 = max((prev2 + num), prev1)
	}

	return prev1
}


