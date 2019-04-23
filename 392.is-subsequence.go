/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 *
 * https://leetcode.com/problems/is-subsequence/description/
 *
 * algorithms
 * Medium (46.55%)
 * Total Accepted:    84.8K
 * Total Submissions: 182.2K
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 * 
 * Given a string s and a string t, check if s is subsequence of t.
 * 
 * 
 * 
 * You may assume that there is only lower case English letters in both s and
 * t. t is potentially a very long (length ~= 500,000) string, and s is a short
 * string (
 * 
 * 
 * A subsequence of a string is a new string which is formed from the original
 * string by deleting some (can be none) of the characters without disturbing
 * the relative positions of the remaining characters. (ie, "ace" is a
 * subsequence of "abcde" while "aec" is not).
 * 
 * 
 * Example 1:
 * s = "abc", t = "ahbgdc"
 * 
 * 
 * Return true.
 * 
 * 
 * Example 2:
 * s = "axc", t = "ahbgdc"
 * 
 * 
 * Return false.
 * 
 * 
 * Follow up:
 * If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you
 * want to check one by one to see if T has its subsequence. In this scenario,
 * how would you change your code?
 * 
 * Credits:Special thanks to @pbrother for adding this problem and creating all
 * test cases.
 */
func isSubsequence2(s string, t string) bool {
	m := len(s)
	n := len(t)
	if m > n {
		return false
	}

	start := 0
	i := 0
	j := 0
	for i = 0; i < m; i++ {
		for j = start; j < n; j++ {
			if s[i] == t[j] {
				start = j + 1
				break
			}
		}

		if j >= n {
			break
		}
	}

	return i > m - 1
}

func isSubsequence(s string, t string) bool {
	m := len(s)
	n := len(t)
	if m > n {
		return false
	}

	i := 0
	j := 0
	for {
		if i >= m || j >= n {
			break
		}

		if s[i] == t[j] {
			i += 1
		}

		j += 1
	}

	return i >= m
}

