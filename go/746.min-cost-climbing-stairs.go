/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 || n == 1 { return 0 }

	for i := 2; i < n; i++ {
		cost[i] += min(cost[i - 1], cost[i - 2])
	}

	return min(cost[n - 1], cost[n - 2])
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

