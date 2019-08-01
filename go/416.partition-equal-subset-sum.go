/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */
 func canPartition(nums []int) bool {
	n := len(nums) 
	if n <= 1 { return false }

	sum := 0
	for _, v := range(nums) {
		sum += v
	}

	if sum % 2 == 1 { return false }
	half := sum / 2

	dp := make([]bool, half + 1)
	dp[0] = true
    for i := 1; i < half + 1; i++ {
        dp[i] = false
    }
    
	for i := 0; i < n; i++ {
		for j := half; j >= 1; j-- {
			if j >= nums[i] {
				dp[j] = dp[j] || dp[j - nums[i]]
			}
		}
	}

    
	return dp[half]
}