/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	state := -1
    for i, v := range word {
		switch(state) {
			case 0, 1:		// all lower case; Capitalize first letter
				if v <= 'Z' {
					return false
				}
			case 2:		// all upper case
				if v >= 'a' {
					return false
				}
			default:
				if i == 0 {
					if v >= 'a' && v <= 'z' {
						state = 0
					}
				} else {
					if v >= 'a' && v <= 'z' {
						state = 1
					} else {
						state = 2
					}
				}
		}
	}

	return true
}
// @lc code=end

