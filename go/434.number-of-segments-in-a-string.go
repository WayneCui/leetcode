/*
 * @lc app=leetcode id=434 lang=golang
 *
 * [434] Number of Segments in a String
 */

// @lc code=start
func countSegments(s string) int {
	state := 0 // 1 - characters state  0 - space state
	count := 0
	for _, c := range s {
		if c == ' ' {
			state = 0
		} else if state == 0 {
			count++
            state = 1
		}
	}

	return count
}
// @lc code=end

