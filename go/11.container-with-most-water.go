/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
 func maxArea(height []int) int {
	water := 0
	i := 0
	j := len(height) - 1

	for i < j {
		water = max(water, min(height[i], height[j]) * (j - i))
		if height[i] < height[j] {
			i += 1
		} else {
			j -= 1
		}
	}

	return water
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

