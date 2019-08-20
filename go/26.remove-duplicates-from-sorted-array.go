/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
func removeDuplicates(nums []int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast += 1
		} else {
			slow += 1
			nums[slow] = nums[fast]
			fast += 1
		}
	}

	return slow + 1
}

