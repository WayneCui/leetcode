/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func maxSubArray(nums []int) int {
	sumSofar := nums[0]
	maxSofar := nums[0]
	for i := 1; i < len(nums); i++ {

		if sumSofar < 0 {
			sumSofar = nums[i]
		} else {
			sumSofar += nums[i]
		}

		if maxSofar < sumSofar {
			maxSofar = sumSofar
		}
	}

	return maxSofar
}

