/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 */

// @lc code=start
func validPalindrome(s string) bool {
	low := 0
	high := len(s) - 1
	for low < high {
		if(s[low:low+1] == s[high:high+1]){
			low += 1
			high -= 1
		} else {
			return isPalindrome(s[low+1:high+1]) || isPalindrome(s[low:high])
		}
	}

	return true
}

func isPalindrome(s string) bool {
	low := 0
	high := len(s) - 1
	for low < high {
		if(s[low:low+1] != s[high:high+1]){
			return false
		}
        
        low, high = low + 1, high - 1
	}

	return true
}
// @lc code=end

