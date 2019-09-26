/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */
func firstUniqChar(s string) int {
	memo := make(map[int]int)
	n := len(s)
	for i := 0; i < n; i++ {
		if _, ok := memo[i]; ok {
			continue
		}

		c := i
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				memo[j] = 1
				c = j
			}
		}
		if(c == i) {
			return i
		}
	}

	return -1
}

func firstUniqChar(s string) int {
	memo := make([]int, 26)
	for _, v := range s {
		memo[v - 'a'] += 1
	}

	for i, v := range s {
		if memo[v - 'a'] == 1 {
			return i
		}
	}

	return -1
}

