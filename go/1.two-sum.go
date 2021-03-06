/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (42.89%)
 * Total Accepted:    1.6M
 * Total Submissions: 3.8M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 * 
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 * 
 * Example:
 * 
 * 
 * Given nums = [2, 7, 11, 15], target = 9,
 * 
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 * 
 * 
 * 
 * 
 */
func twoSum0(nums []int, target int) []int {
	var length = len(nums)
	dict := make(map[int] int)

	for i := 0; i < length; i++ {
		dict[nums[i]] = i
	}

	for i := 0; i < length; i++ {
		first := nums[i]
		second := target - first
		idx, ok := dict[second]

		if ok && i != idx {
			return []int{i, dict[second]}
		}
	}

	return []int{-1, -1}
}
// a more succinct version
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if idx, ok := dict[target - nums[i]]; ok {
			return []int{idx, i}
		} else {
			dict[nums[i]] = i
		}
	}

	return []int{-1, -1}
}

