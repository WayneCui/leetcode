/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */
 func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	memo := make([]int, m)
	memo[0] = grid[0][0]
	for i := 1; i < m; i++ {
		memo[i] = memo[i - 1] + grid[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				memo[j] = memo[j] + grid[i][j]
			} else {
				memo[j] = min(memo[j], memo[j - 1]) + grid[i][j]
			}
			
		}
	}

	return memo[m - 1]
 }

 func min(a, b int ) int {
	 if a < b {
		 return a
	 } else {
		 return b
	 }
 }

