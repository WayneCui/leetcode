/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}

	sum := make([]int, n)
	for i := 0; i < n; i++ {
		sum[i] = triangle[n - 1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j = 0; j < len(triangle[i]); j++ {
			sum[j] = triangle[i][j] + min(sum[j], sum[j + 1])
		}
	}

	return sum[0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

