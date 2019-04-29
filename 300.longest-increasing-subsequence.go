/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	memo := make(map[int]int)
	for i := 0; i < n; i++ {
		memo[i] = 0
	}

	memo[0] = 1
	count := 1
	for i := 1; i < n; i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				maxval = max(maxval, memo[j])
			}
		}

		memo[i] = maxval + 1
		count = max(count, memo[i])
	}

	return count
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	memo := []int{}
	for i := 0; i < n; i++ {
		memo = append(memo, 0)
	}

	size := 0
	for i := 0; i < n; i++ {
		idx := binarySearch(memo, 0, size, nums[i])
		if idx < 0 {
			idx = -idx - 1
		}

		memo[idx] = nums[i]
		
		if idx == size {
			size += 1
		}
	}

	return size
}

func binarySearch(nums []int, from, to, target int) int {
	low := from
	high := to - 1
	for low <= high {
		mid := (low + high) / 2
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -(low + 1)
}
