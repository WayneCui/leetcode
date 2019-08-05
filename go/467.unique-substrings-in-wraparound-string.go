/*
 * @lc app=leetcode id=467 lang=golang
 *
 * [467] Unique Substrings in Wraparound String
 */
 func findSubstringInWraproundString0(p string) int {
	n := len(p)
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	parts := []int{0}
	for i := 1; i < n; i++ {
		if p[i] == p[i-1] + 1 || (p[i] == 97 && p[i - 1] == 122) {
			continue
		}
        
        parts = append(parts, i)
	}

	parts = append(parts, n)
	memo := make(map[string]int)
	for i := 1; i < len(parts); i++ {
		collect(parts[i] - parts[i - 1], int(p[parts[i - 1]]), memo)
	}
	return len(memo)
}

func collect(length int, from int, memo map[string]int)  {
	k := string(97 + (from - 97) % 26) + strconv.Itoa(length)
	if _, ok := memo[k]; ok {
		return
	}
	for i := 0; i < length; i++ {
        if i >= 26 { return }
		for j := 1; j <= length - i; j++ {
			key := string(97 + (from + i - 97) % 26) + strconv.Itoa(j)
			memo[key] = 1
		}
	}
}

//credits to https://leetcode.com/problems/unique-substrings-in-wraparound-string/discuss/95439/Concise-Java-solution-using-DP
func findSubstringInWraproundString(p string) int {
	n := len(p)
	if n == 0 { return 0 }

	dp := make([]int, 26)

	maxSofar := 1
    dp[p[0] - 97] = 1
	for i := 1; i < n; i++ {
		if p[i] == p[i - 1] + 1 || p[i-1] - p[i] == 25 {
			maxSofar++
		} else {
			maxSofar = 1
		}

		if maxSofar > dp[p[i] - 97] {
			dp[p[i] - 97] = maxSofar
		}
	}

	sum := 0
	for _, v := range(dp) {
		sum += v
	}

	return sum
}