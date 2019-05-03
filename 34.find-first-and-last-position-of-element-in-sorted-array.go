/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (33.30%)
 * Total Accepted:    286K
 * Total Submissions: 859K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * If the target is not found in the array, return [-1, -1].
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 * 
 */
func searchRange1(nums []int, target int) []int {
	n := len(nums)
	result := []int{-1, -1}

	if n == 0 {
		return result
	}

	low := 0
	high := n - 1
	
	for low < high {

		if target == nums[low] {
			result[0] = low
		} 

		if target == nums[high] {
			result[1] = high
		}

		if result[0] >= 0 && result[1] >= 0 {
			break
		}

		mid := (low + high) / 2
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid
		} else {
			if result[0] < 0 {
				low += 1
			} 
			
			if result[1] < 0 {
				high -= 1
			}
		}
	}

	if target == nums[low] {
		if result[0] < 0 {
			result[0] = low
		}

		if result[1] < 0 {
			result[1] = low
		}
	}

	return result
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	result := []int{-1, -1}

	if n == 0 {
		return result
	}

	low := 0
	high := n - 1
	
	for low < high {

		if target == nums[low] {
			result[0] = low
		} 

		if target == nums[high] {
			result[1] = high
		}

		if result[0] >= 0 && result[1] >= 0 {
			break
		}

		mid := (low + high) / 2
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid
		} else {
			low1 := searchLeft(nums, target, low, mid)
			high1 := searchRight(nums, target, mid, high)
			result[0] = low1
			result[1] = high1
		}
	}

	if target == nums[low] {
		if result[0] < 0 {
			result[0] = low
		}

		if result[1] < 0 {
			result[1] = low
		}
	}

	return result
}

func searchLeft(nums []int, target int, from int, to int) int {
	for from < to {
		mid := (from + to) / 2
		if target > nums[mid] {
			from = mid + 1
		} else {
			to = mid
		}
	}

	return to
}

func searchRight(nums []int, target int, from int, to int) int {
	for from < to {
		mid := (from + to) / 2
		if target == nums[mid] {
			from = mid + 1
		} else {
			to = mid
		}
	}

	return from - 1
}

