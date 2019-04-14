/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (32.75%)
 * Total Accepted:    393.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 * 
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 * 
 * You may assume no duplicate exists in the array.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * 
 */
 func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	ok := false
	low := 0

	pivotIdx, pivot := findMin(nums)
	if pivotIdx == 0 {
		ok, low = binarySearch(nums, 0, n - 1, target)
	} else if target >= pivot && target <= nums[n - 1] {
		ok, low = binarySearch(nums, pivotIdx, n - 1, target)
	} else {
		ok, low = binarySearch(nums, 0, pivotIdx - 1, target)
	}

	if ok {
		return low
	} else {
		return -1
	}
}

func findMin(nums []int) (int, int) {
	n := len(nums)
	if nums[0] < nums[n - 1] {
		return 0, nums[0]
	}

	low := 0
	high := n - 1

	for low < high {
		if nums[low] < nums[high] {
			break
		} else {
			mid := (low + high) / 2
			if nums[mid] >= nums[0] {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}

	return low, nums[low]
}

func binarySearch(nums []int, low int, high int, target int)(bool, int) {
	for low < high {
		mid := (low + high) / 2
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return target == nums[low], low
}

