/*
 * @lc app=leetcode id=686 lang=golang
 *
 * [686] Repeated String Match
 */

// @lc code=start
func repeatedStringMatch(A string, B string) int {
    times := int(math.Ceil(float64(len(B)) / float64(len(A))))

	i := 0
	for i < 2 {
		repeated := strings.Repeat(A, times + i)
		if(strings.Contains(repeated, B)) {
			return times + i
		}
		i += 1
	}
	
	return -1
}
// @lc code=end

