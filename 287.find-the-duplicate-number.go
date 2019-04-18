/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 *
 * https://leetcode.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (49.03%)
 * Total Accepted:    179.2K
 * Total Submissions: 365.3K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * Given an array nums containing n + 1 integers where each integer is between
 * 1 and n (inclusive), prove that at least one duplicate number must exist.
 * Assume that there is only one duplicate number, find the duplicate one.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,3,4,2,2]
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,1,3,4,2]
 * Output: 3
 * 
 * Note:
 * 
 * 
 * You must not modify the array (assume the array is read only).
 * You must use only constant, O(1) extra space.
 * Your runtime complexity should be less than O(n^2).
 * There is only one duplicate number in the array, but it could be repeated
 * more than once.
 * 
 * 
 */
// func findDuplicate(nums []int) int {
// 	m := len(nums)
// 	for i := 0; i < m - 1; i++ {
// 		for j := i + 1; j < m; j++ {
// 			if nums[j] == nums[i] {
// 				return nums[i]
// 			}
// 		}
// 	}

// 	return 0
// }


func findDuplicate(nums []int) int {
	if len(nums) > 0 {
		slow := nums[0]
		fast := nums[nums[0]]

		for {
			slow = nums[slow]
			fast = nums[nums[fast]]

			if slow == fast {
				break;
			}
		}

		finder := 0
		for {
			slow = nums[slow]
			finder = nums[finder]

			if slow == finder {
				return slow
			}
		}
	}

	return -1
}

