/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */
func integerBreak(n int) int {
	if n == 2 { return 1 }
	if n == 3 { return 2 }

	k := float64(n / 3)
	r := n % 3
	result := 0.0
	if r == 2 {
		result = math.Pow(3, k) * 2
	} else if r == 1 {
		result = math.Pow(3, k - 1) * 4
	} else {
		result = math.Pow(3, k)
	}

	return int(result)
}

