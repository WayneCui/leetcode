/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

//Time Limit Exceeded
func wordBreak0(s string, wordDict []string) bool {
	if len(s) == 0 { return false }

	return toBreak(s, wordDict)
}

func toBreak(s string, wordDict []string) bool {
	if len(s) == 0 { return true }

	result := false
	for _, w := range(wordDict) {
		if strings.hasPrefix(s, w) {
			result = toBreak(s[len(w):], wordDict)
			if result {
				return true
			}
		}
	}

	return result
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	if n == 0 { return false }
    
	dict := make(map[string]int)
	for _, v := range(wordDict) {
		dict[v] = 1
	}

    dp := make([]bool, n)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			_, ok := dict[s[j : i + 1]]
			if (j == 0  || dp[j - 1] ) && ok {
				dp[i] = true
			}
		}
	}
    //fmt.Println(dp)

	return dp[n - 1]
}