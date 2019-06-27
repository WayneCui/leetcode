/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
func climbStairs0(n int) int {
	if n == 0 { return 1 }
	if n == 1 { return 1 }

	memo := make(map[int]int)
	memo[0] = 1
	memo[1] = 1
	return climbStairsWithMemo(n, memo)
}

func climbStairsWithMemo(n int, memo map[int]int) int {
	v, ok := memo[n]
	if ok {
		return v
	} else {
		v := climbStairsWithMemo(n - 1, memo) + climbStairsWithMemo(n - 2, memo)
		memo[n] = v
		return v
	}
}

func climbStairs1(n int) int {
	if n == 0 { return 1 }
	if n == 1 { return 1 }

	memo := make(map[int]int)
	memo[0] = 1
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i - 1] + memo[i - 2]
	}

	return memo[n]
}

func climbStairs(n int) int {
	if n == 0 { return 1 }
	if n == 1 { return 1 }

	twoStepBefore := 0
	oneStepBefore := 1
	for i := 1; i <= n; i++ {
		oneStepBefore, twoStepBefore = oneStepBefore + twoStepBefore, oneStepBefore
	}

	return oneStepBefore
}



