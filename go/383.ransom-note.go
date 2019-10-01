/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	memo := [26]int{0}
	for _, letter := range magazine {
		memo[letter - 'a']++ 
	}

	for _, letter := range ransomNote {
		if memo[letter - 'a'] <= 0 {
			return false
		} else {
			memo[letter - 'a']--
		}
	}

	return true
}
// @lc code=end

