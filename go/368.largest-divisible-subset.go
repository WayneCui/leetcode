/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 */

 //Memory Limit Exceeded	
 func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)
	optional := [][]int{}
    optional = append(optional, []int{nums[0]})

	cur := optional[0] 
	curLen := 1
	for i := 1; i < len(nums); i++ {
        n := len(optional)
		for j := 0; j < n; j++ {
			if nums[i] % optional[j][len(optional[j]) - 1] == 0 {
				optional[j] = append(optional[j], nums[i])
				if len(optional[j]) > curLen {
					curLen = len(optional[j])
					cur = optional[j]
				}
			} else {
				optional = append(optional, []int{ nums[i]} )
			}
		}
	}

    return cur
}

// credits to https://leetcode.com/problems/largest-divisible-subset/discuss/84006/Classic-DP-solution-similar-to-LIS-O(n2)
func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	sort.Ints(nums)
	count := make([]int, n)
	pre := make([]int, n)

	max := 0
	idx := -1
	for i := 0; i < n; i++ {
		count[i] = 1
		pre[i] = -1
		for j := i - 1; j >= 0; j-- {
			if nums[i] % nums[j] == 0 && count[j] + 1 > count[i] {
				count[i] = count[j] + 1
				pre[i] = j
			}
		}

		if max < count[i] {
			max = count[i]
			idx = i
		}
	}

	result := []int {}
    fmt.Println(pre)
	for idx > -1 {
		result = append(result, nums[idx])
		idx = pre[idx]
	}

	return result
}

