/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */
func compress(chars []byte) int {
	count := 1
	b := chars[0]
	i := 0
	curr := 0
	for i < len(chars){
		if chars[i] == b {
			count++
		} else {
			chars[curr] = b
			curr++
			if count > 1 {
				for _, v := range strconv.Itoa(count) {
					chars[curr] = byte(v)
					curr++
				}
			}

			b = chars[i]
			count = 1
			i++
		}
	}

	chars[curr] = b
	curr++
	if count > 1 {
		for _, v := range strconv.Itoa(count) {
			chars[curr] = byte(v)
			curr++
		}
	}

	return buf.Len()
}

