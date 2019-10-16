/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 */

// @lc code=start
func checkRecord(s string) bool {
	A := 0
	L1 := 0
	prev := ""
    prevprev := ""
    for _, v := range s {
		switch string(v) {
			case "A":
				A += 1
			case "L":
				if prev == "L" && prevprev == "L" {
					L1 += 1
				}
			default:
		}
        prevprev = prev
        prev = string(v)
	}

    if A > 1 || L1 >= 1 {
		return false
	}
	return true
}
// @lc code=end

