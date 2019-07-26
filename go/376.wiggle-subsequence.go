/*
 * @lc app=leetcode id=376 lang=golang
 *
 * [376] Wiggle Subsequence
 */
 func wiggleMaxLength0(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	prev := nums[0]
	val := -1

	k := 1
	flag := 0
	for k = 1; k < n; k++ {
		if nums[k] - prev < 0 {
			flag = -1
            break
		} else if nums[k] - prev > 0 {
			flag = 1
            break
		}
	}
	
	count := 1
	for i := k; i < n; i++ {
		val = nums[i]
		if flag * (val - prev) > 0 {
			flag = 0 - flag
			count += 1
		}

		prev = val
	}

	return count
}

func wiggleMaxLength1(nums []int) int {
	if len(nums) == 0 { return 0 }

	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i - 1] {
			up = down + 1
		} else if nums[i] < nums[i - 1] {
			down = up + 1
		}
	}

	return max(up, down)
}

func max(a, b int ) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }

	up, down := make([]int, n), make([]int, n)
	up[0] = 1
	down[0] = 1

	for i := 1; i < n; i++ {
		if nums[i] > nums[i - 1] {
			up[i] = down[i - 1] + 1
			down[i] = down[i - 1]
		} else if nums[i] < nums[i - 1] {
			down[i] = up[i - 1] + 1
			up[i] = up[i - 1]
		} else {
			down[i] = down[i - 1]
			up[i] = up[i - 1]
		}
	}

	return max(up[n - 1], down[n - 1])
}