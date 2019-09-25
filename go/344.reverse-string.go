/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */
func reverseString(s []byte)  {
	n := len(s)
	for i := 0; i < n / 2; i++ {
		s[i], s[ n - 1 - i]= s[ n - 1 - i], s[i]
	}
}

