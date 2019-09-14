/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
 func generateParenthesis(n int) []string {
    res := []string{}
	var buf bytes.Buffer
	addParen(&res, n, n, buf)
	return res
}

func addParen(collector *[]string, lfRemain int, rgRemain int, buf bytes.Buffer) {
	if lfRemain < 0 || rgRemain < lfRemain {
		return
	}

	if lfRemain == 0 && rgRemain == 0 {
		*collector = append(*collector, buf.String())
		return
	} 

	if lfRemain > 0 {
		newbuf := bytes.NewBufferString(buf.String())
		newbuf.WriteString("(")
		addParen(collector, lfRemain - 1, rgRemain, *newbuf)
	}

	if rgRemain > lfRemain {
		newbuf := bytes.NewBufferString(buf.String())
		newbuf.WriteString(")")
		addParen(collector, lfRemain, rgRemain - 1, *newbuf)
	}
}


