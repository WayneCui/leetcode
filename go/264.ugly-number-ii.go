/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */
func nthUglyNumber(n int) int {
	if n == 1 { return 1 }
	memo := make([]int, n)
	memo[0] = 1

	p2 := 0
	p3 := 0
	p5 := 0
	for i := 1; i < n; i++ {
		memo[i] = min(memo[p2] * 2, memo[p3] * 3, memo[p5] * 5)
		if memo[i] == memo[p2] * 2 { p2 += 1 }
		if memo[i] == memo[p3] * 3 { p3 += 1 }
		if memo[i] == memo[p5] * 5 { p5 += 1 }
	}

	return memo[n - 1]
}

func min(a ...int) int {
	m := a[0]
	for _, v := range(a) {
		if v < m {
			m = v
		}
	}

	return m
}

