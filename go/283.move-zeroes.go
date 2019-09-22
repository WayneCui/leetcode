/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */
func moveZeroes(nums []int)  {
	 n := len(nums)
	 for i := 0; i < n - 1; i++ {
		 if nums[i] == 0 {
			for j := i + 1; j < n; j++ {
				if nums[j] != 0 {
					tmp := nums[i]
					nums[i] = nums[j]
					nums[j] = tmp
					break
				}
			}
		 }
	 }
}

