/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */
 func maxProfit(prices []int) int {
	n := len(prices)
	if(n < 2) {
		return 0
	}
	buyMemo := make([]int, n)
	sellMemo := make([]int, n)
	buyMemo[0] = -prices[0]
    buyMemo[1] = max(-prices[0], -prices[1])
	sellMemo[0] = 0
    sellMemo[1] = max(0, prices[1] - prices[0])

	for i := 2; i < n; i++ {
		buyMemo[i] = max(buyMemo[i - 1], sellMemo[i - 2] - prices[i])
		sellMemo[i] = max(sellMemo[i - 1], buyMemo[i - 1] + prices[i])
	}
    
    return sellMemo[n - 1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

