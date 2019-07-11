/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }
	if n == 1 { return nums[0] }

	mx := nums[0]
	mn := nums[0]
	maxSofar := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			tmp := mx
			mx = mn
			mn = tmp
		}
		mx = max(nums[i], mx * nums[i])
		mn = min(nums[i], mn * nums[i])

		maxSofar = max(maxSofar, mx)
	}

	return maxSofar
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
