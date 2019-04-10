/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (40.59%)
 * Total Accepted:    381K
 * Total Submissions: 937.7K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array and a target value, return the index if the target is
 * found. If not, return the index where it would be if it were inserted in
 * order.
 * 
 * You may assume no duplicates in the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,3,5,6], 5
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [1,3,5,6], 2
 * Output: 1
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: [1,3,5,6], 7
 * Output: 4
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: [1,3,5,6], 0
 * Output: 0
 * 
 * 
 */
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := 0

	for low <= high {
		mid = (high - low ) / 2 + low
		midVal := nums[mid]

		switch {
		case midVal < target: 
			low = mid + 1
		case midVal > target:
			high = mid - 1
		default:
			return mid
		}
	}

	return low
}

