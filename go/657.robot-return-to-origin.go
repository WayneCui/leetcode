/*
 * @lc app=leetcode id=657 lang=golang
 *
 * [657] Robot Return to Origin
 */

// @lc code=start
func judgeCircle(moves string) bool {
	bytes := []byte(moves)
	x, y := 0, 0
	for i := 0; i < len(bytes); i++ {
		switch bytes[i] {
		case 'L':
			x -= 1
		case 'R':
			x += 1
		case 'U':
			y += 1
		case 'D':
			y -= 1
		default:
		}
	}

	return x == 0 && y == 0
}

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, v := range moves {
		switch v {
		case 'L':
			x -= 1
		case 'R':
			x += 1
		case 'U':
			y += 1
		case 'D':
			y -= 1
		default:
		}
	}

	return x == 0 && y == 0
}
// @lc code=end

