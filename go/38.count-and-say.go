/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */
func countAndSay(n int) string {
	str := "1"
	for i := 0; i < n - 1; i++ {
		str = helper(str)
	}

	return str
}

func helper(s string) string {
	i := 1
	c := 1
	curr := s[0]
	var buf bytes.Buffer
	for i < len(s) {
		if curr == s[i] {
			c += 1
		} else {
			buf.WriteString(strconv.Itoa(c))
			buf.WriteString(string(curr))
			curr = s[i]
			c = 1
		}
	}
	buf.WriteString(strconv.Itoa(c))
	buf.WriteString(string(curr))

	fmt.Println(buf.String())
	return buf.String()
}
