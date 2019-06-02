/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */
func numTrees(n int) int {
	if n == 0 { return 1 }
	if n == 1 { return 1 }

	result := 0
	for i := 1; i <= n / 2; i++ {
		result += numTrees(i - 1) * numTrees(n - i)
	}
	result = result * 2

	if n % 2 == 1 {
		i := n / 2 + 1
		result += numTrees(i - 1) * numTrees(n - i)
	}

	return result
}

