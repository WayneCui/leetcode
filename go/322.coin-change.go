/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */
 func coinChange0(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, amount + 1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			for j, _ := range(coins) {
				dp[i][j] = 0
			}
		} else {
			for j, _ := range(coins) {
				dp[i][j] = math.MaxInt64
			}
		}
	}

    for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			if i > coins[j] {
				if min(dp[ i - coins[j]]) == math.MaxInt64 {
					dp[i][j] = math.MaxInt64
				} else {
                    dp[i][j] = 1 + min(dp[ i - coins[j]])
				}
			} else if i == coins[j] {
				dp[i][j] = 1
			}
		}
	}

	// fmt.Println(dp)
	if min(dp[amount]) == math.MaxInt64 {
		return -1
	} else {
		return min(dp[amount])
	}
}

func min(memo []int) int {
	least := math.MaxInt64
	for _, v := range(memo) {
		if v < least {
			least = v
		}
	}

	return least
}

func coinChange1(coins []int, amount int) int {
	n := len(coins)
	memo := make([]int, amount + 1)
	memo[0] = 0
	for i := 1; i <= amount; i++ {
		memo[i] = math.MaxInt64
	}

    for i := 1; i <= amount; i++ {
		least := math.MaxInt64
		for j := 0; j < n; j++ {
			if i > coins[j] {
				if memo[ i - coins[j]] == math.MaxInt64 {
					if least > math.MaxInt64 {
						least = math.MaxInt64
					}
				} else {
					if least > 1 + memo[ i - coins[j]] {
						least = 1 + memo[ i - coins[j]]
					}
				}
			} else if i == coins[j] {
				least = 1
			}
		}

		memo[i] = least
	}

	//fmt.Println(memo)
	if memo[amount] == math.MaxInt64 {
		return -1
	} else {
		return memo[amount]
	}
}

//too slow, 988 ms	
func coinChange2(coins []int, amount int) int {
    if amount == 0 {
        return 0
	}
	
	layer := make(map[int]int)
	layer[amount] = 1

	count := 0
	outer:
    for len(layer) > 0 {
        //fmt.Println(layer)
		nextLayer := make(map[int]int)

		for k := range(layer) {
			for _, c := range(coins) {
				if k - c > 0 {
					nextLayer[k - c] = 1
				} else if k - c == 0 {
					break outer
				}	
			}
		}

		count += 1
		layer = nextLayer
	}

	if len(layer) == 0 { return -1 } else { return count + 1 }
}

//too slow, 240 ms	
func coinChange3(coins []int, amount int) int {
    if amount == 0 {
        return 0
	}
	
	seen := make(map[int]int)
	seen[amount] = 1
	layer := make(map[int]int)
	layer[amount] = 1

	count := 0
	outer:
    for len(layer) > 0 {
        //fmt.Println(layer)
		nextLayer := make(map[int]int)

		for k := range(layer) {
			for _, c := range(coins) {
				_, ok := seen[k - c]
				if k - c > 0 {
					if !ok {
						nextLayer[k - c] = 1
						seen[k - c] = 1
					}
				} else if k - c == 0 {
					break outer
				}	
			}
		}

		count += 1
		layer = nextLayer
	}

	if len(layer) == 0 { return -1 } else { return count + 1 }
}

func coinChange(coins []int, amount int) int {
	n := len(coins)
	memo := make([]int, amount + 1)
	memo[0] = 0
	for i := 1; i <= amount; i++ {
		memo[i] = amount + 1
	}

    for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			if i >= coins[j] {
				if memo[i] > 1 + memo[ i - coins[j]] {
					memo[i] = 1 + memo[ i - coins[j]]
				}
			} 
		}
	}

	//fmt.Println(memo)
	if memo[amount] > amount {
		return -1
	} else {
		return memo[amount]
	}
}

