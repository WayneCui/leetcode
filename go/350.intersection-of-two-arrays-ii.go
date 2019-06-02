/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */
func intersect1(nums1 []int, nums2 []int) []int {
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

func intersect(nums1 []int, nums2 []int) []int {
	m := len(nums1)
	n := len(nums2)
	var out []int
	cache := make(map[int]int)
	for i := 0; i < m; i++ {
		_, ok := cache[nums1[i]]
		if ok {
			cache[nums1[i]] = cache[nums1[i]] + 1
		} else {
			cache[nums1[i]] = 1
		}
	}

	for j := 0; j < n; j++ {
		count, ok := cache[nums2[j]]
		if ok && count > 0{
			out = append(out, nums2[j])
			cache[nums2[j]] = cache[nums2[j]] - 1
		}
	}

	return out
}

