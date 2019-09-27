/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
func romanToInt(s string) int {
	n := len(s)
	result := 0
    for i := 0; i < n; i++ {
		switch string(s[i]) {
		case "I":
			if i < n - 1 && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
				result -= 1
			} else {
				result += 1
			}
		case "V":
			result += 5
		case "X":
			if i < n - 1 && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
				result -= 10
			} else {
				result += 10
			}
		case "L":
			result += 50
		case "C":
			if i < n - 1 && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
				result -= 100
			} else {
				result += 100
			}
		case "D":
			result += 500
		case "M":
			result += 1000
		default:
		}
	}

	return result
}