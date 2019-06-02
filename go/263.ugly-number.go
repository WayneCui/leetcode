/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */
func isUgly(num int) bool {
	if num == 0 {
		return false
	}

	if num == 1 || num == 2 || num == 3 || num == 5 { return true }

	for {
		if num % 2 == 0 {
			num = num / 2
		} else if num % 3 == 0 {
			num = num / 3
		} else if num % 5 == 0 {
			num = num / 5
		} else {
			break
		}
	}

	if num == 1 { 
		return true
	} else {
		return false
	}
}

