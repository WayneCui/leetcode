/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */
//1444ms
func numSquares0(n int) int {
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	return cal(n, memo)
}

func cal(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	count := n + 1
    for i := 1; i <= n / 2 + 1; i++ {
		square := i * i
		if n >= square {
			_c := 1 + cal(n - square, memo)
			if count > _c {
				count = _c
			}
		}
	}
	memo[n] = count
	return count
}


func numSquares2(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	for i := 1; i <= n; i++ {
		//fmt.Println(dp)
		for j := i / 2 + 1; j >= 1; j-- {
			square := j * j
			if i - square >= 0 {
				if dp[i] > 1 + dp[i - square] {
					dp[i] = 1 + dp[i - square]
				}
			}
		}
	}

	return dp[n]
}

func numSquares(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	for i := 1; i <= n; i++ {
		//fmt.Println(dp)
		for j := i / 2 + 1; j >= 1; j-- {
			square := j * j
			if i - square >= 0 {
				if dp[i] > 1 + dp[i - square] {
					dp[i] = 1 + dp[i - square]
				}
			}
		}
	}

	return dp[n]
}

