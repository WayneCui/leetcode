/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (30.94%)
 * Total Accepted:    348.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 * 
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 * 
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 * 
 * Example 1:
 * 
 * 
 * Input: 4
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since 
 * the decimal part is truncated, 2 is returned.
 * 
 * 
 */
func mySqrt(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}

	low := 0.0
	high := float64(x)
	mid := 0.0
	square := 0
	
	precise := 1e-7
	for high - low > precise {
		mid = (low + high) / 2
		square = int(mid * mid)
		if square > x {
			high = mid
		} else if square < x {
			low = mid
		} else {
			break
		}
	}

	return int((high + low) / 2)
}
