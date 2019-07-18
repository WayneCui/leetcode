/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */
func countBits(num int) []int {
	if num == 0 { return []int{0} }
	if num == 1 { return []int{0, 1} }

	dp := make([]int, num + 1)
	dp[0] = 0
	flag := 1
	for n := 1; n <= num; n++ {
		if n % flag == 0 {
			dp[n] = 1
			flag = n
		} else {
			dp[n] = dp[n - flag] + 1
		}
	}

	return dp
}

