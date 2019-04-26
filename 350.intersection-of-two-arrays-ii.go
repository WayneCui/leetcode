/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */
func intersect(nums1 []int, nums2 []int) []int {
	m := len(nums1)
	n := len(nums2)
	var out []int
	cache := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			_, ok := cache[j]
			if ok { continue }
			if nums1[i] == nums2[j] {
				out = append(out, nums1[i])
				cache[j] = 1
				break
			}
		}
	}

	return out
}

