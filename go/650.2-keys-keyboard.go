/*
 * @lc app=leetcode id=650 lang=golang
 *
 * [650] 2 Keys Keyboard
 */
 func minSteps(n int) int {
	if n == 1 { return 0 }
	memo := make([]int, n + 1)
	for i := 1; i < n + 1; i++ {
		memo[i] = i
	}

    for i := 2; i <= n; i++ {
		if n % i != 0 { continue }
		for j := 2; j < i / 2 + 1; j++ {
			if i % j == 0 {
				memo[i] = min(memo[i], memo[j] + (i / j))
			}
		}
	}

	return memo[n]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
