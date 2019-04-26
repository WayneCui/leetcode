/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */
func intersection(nums1 []int, nums2 []int) []int {
	m := len(nums1)
	n := len(nums2)
	memo := make(map[int]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] == nums2[j] {
				memo[nums1[i]] = 0
				break
			}
		}
	}

	var keys []int
	for k := range memo {
		keys = append(keys, k)
	}
	return keys
}

