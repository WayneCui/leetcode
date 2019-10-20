/*
 * @lc app=leetcode id=521 lang=golang
 *
 * [521] Longest Uncommon Subsequence I 
 */

// @lc code=start
func findLUSlength(a string, b string) int {
    n1 := len(a)
    n2 := len(b)
    if n1 > n2 {
        return n1
    } else if n2 > n1 {
        return n2
    } else {
        if a == b {
            return -1
        } else {
            return n1
        }
    }
}
// @lc code=end

