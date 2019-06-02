/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (25.89%)
 * Total Accepted:    407.7K
 * Total Submissions: 1.6M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * 
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 * 
 * You may assume nums1 and nums2 cannot be both empty.
 * 
 * Example 1:
 * 
 * 
 * nums1 = [1, 3]
 * nums2 = [2]
 * 
 * The median is 2.0
 * 
 * 
 * Example 2:
 * 
 * 
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * 
 * The median is (2 + 3)/2 = 2.5
 * 
 * 
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	r := (m + n) % 2
	q := (m + n) / 2

	i1 := 0
	i2 := 0

	a := 0
	b := 0
	for i := 0; i < m + n; i++ {
		if i1 < m && i2 < n {
			if nums1[i1] <= nums2[i2] {
				b = a
				a = nums1[i1]
				i1 = i1 + 1
			} else {
				b = a
				a = nums2[i2]
				i2 = i2 + 1
			}
		} else if i1 < m {
			b = a
			a = nums1[i1]
			i1 = i1 + 1
		} else if i2 < n {
			b = a
			a = nums2[i2]
			i2 = i2 + 1
		}

		if i == q {
			break
		}
	}

	if r == 0 {
		return (float64(a) + float64(b)) / 2.0
	} else {
		return float64(a)
	}
}

