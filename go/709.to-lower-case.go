/*
 * @lc app=leetcode id=709 lang=golang
 *
 * [709] To Lower Case
 */

// @lc code=start
func toLowerCase(str string) string {
	runes := []rune(str)
    for i, c := range str {
		if c >= 'A' && c <= 'Z' {
			runes[i] = c + ('a' - 'A')
		}
	}

	return string(runes)
}
// @lc code=end

