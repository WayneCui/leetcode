/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

//Time Limit Exceeded [2,1,3] 35
func combinationSum4_0(nums []int, target int) int {
	count := 0 
    for _, v := range(nums) {
		if target - v == 0 {
			count++
		} else if target - v > 0 {
			count += combinationSum4(nums, target - v)
		}
	}

	return count
}

func combinationSum4(nums []int, target int) int {
	memo := make(map[int]int)
	return doCombination(nums, target, memo)
}

func doCombination(nums []int, target int, memo map[int]int) int {
	if v, ok := memo[target]; ok {
		return v
	}

	count := 0 
    for _, v := range(nums) {
		if target - v == 0 {
			count++
		} else if target - v > 0 {
			count += doCombination(nums, target - v)
		}
	}

	memo[target] = count
	return count
}
