/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
 func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 { return "" }
	output := ""
	idx := 0
	currChar := ""
	cntl := true
	for cntl {
        output += currChar
        currChar = ""
        
		for _, s := range(strs) {
			if idx < len(s) {
				if currChar == "" {
					currChar = string(s[idx])
				} else {
					if string(s[idx]) != currChar {
                        cntl = false
					} 
				}
			} else {
				cntl = false
			}
		}

		idx += 1
	}

	return output
}

