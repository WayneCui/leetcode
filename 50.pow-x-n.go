/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (27.81%)
 * Total Accepted:    308.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 * 
 * Example 1:
 * 
 * 
 * Input: 2.00000, 10
 * Output: 1024.00000
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 2.10000, 3
 * Output: 9.26100
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 * 
 * 
 * Note:
 * 
 * 
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 * 
 * 
 */
func myPow(x float64, n int) float64 {
	switch {
		case x == 1.0:
			return 1.0
		case x == -1.0:
			if n % 2 == 0 {
				return 1.0
			} else {
				return -1.0
			}
		case n == 0:
			return 1
		case n == 1:
			return x
		case n < 0:
			return 1 / positivePow(x, -n)
		default:
			return positivePow(x, n)
	}
}

func positivePow(x float64, n int) float64 {
	product := 1.0

	for i := 0; i < n; i++ {
		product *= x
	}

	return product
}

