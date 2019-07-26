/*
 * @lc app=leetcode id=375 lang=golang
 *
 * [375] Guess Number Higher or Lower II
 */
func getMoneyAmount0(n int) int {
	memo := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n + 1)
	}

	return guess(1, n, memo)
}

func guess(from int, to int, memo [][]int ) int {
	if to <= from { return 0 }
	if memo[from][to] != 0 { return memo[from][to] }

	res := math.MaxInt64
	for i := from; i < to; i++ {
		tmp := i + max(guess(from, i - 1, memo), guess(i + 1, to, memo))
		if tmp < res {
			res = tmp
		}
	}

	memo[from][to] = res

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n + 1)
	for i := 0; i < n + 1; i++ {
		dp[i] = make([]int, n + 1)
	}

	for lo := n; lo >= 1; lo-- {
		for hi := lo + 1; hi < n + 1; hi++ {
			res := math.MaxInt64
			for i := lo; i < hi; i++ {
				tmp := i + max(dp[lo][i - 1], dp[i + 1][hi])
				if tmp < res {
					res = tmp
				}
			}
			dp[lo][hi] = res
		}
	}

	return dp[1][n]
}
