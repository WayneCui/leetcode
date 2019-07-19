/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */
func countNumbersWithUniqueDigits(n int) int {
	total := 1
	prev := 0
	num := 0
	for k := 1; k <= n; k++ {
		if k == 1 {
			num = 9
		} else {
			num = prev * (10 - k + 1)
		}
		prev = num
		total += num
	}

	return total
}

