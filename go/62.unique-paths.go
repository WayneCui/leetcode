/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
//Time Limit Exceeded
func uniquePaths0(m int, n int) int {
	if m == 1 || n == 1 { return 1 }
	if m == 2 && n == 2 { return 2 }

	return uniquePaths(m - 1, n) + uniquePaths(m, n - 1)
}

//Time Limit Exceeded
func uniquePaths1(m int, n int) int {
	var memo map[string]int
	return uniquePathsWithMemo(m, n, memo)
}

func uniquePathsWithMemo(m int, n int, memo map[string]int) int {
	if m == 1 || n == 1 { return 1 }
	if m == 2 && n == 2 { return 2 }

	k := key(m, n)
	v, ok := memo[k]
	if ok {
		return v
	}

	return uniquePathsWithMemo(m - 1, n, memo) + uniquePathsWithMemo(m, n - 1, memo)
}

func uniquePaths2(m int, n int) int {
	if m == 1 || n == 1 { return 1 }
	if m == 2 && n == 2 { return 2 }

	memo := make(map[string]int)
    
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || j == 1 {
				memo[key(i, j)] = 1
			} else {
				memo[key(i, j)] = memo[key(i - 1, j)] + memo[key(i, j - 1)]
			}
		}
	}
    //fmt.Println(memo)
	
	return memo[key(m, n)]
}

func key(m, n int) string {
	if m < n {
		return strconv.Itoa(m) + ":" + strconv.Itoa(n)
	} else {
		return strconv.Itoa(n) + ":" + strconv.Itoa(m)
	}
}

func uniquePaths3(m int, n int) int {
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				memo[i][j] = 1
			} else {
				memo[i][j] = memo[i - 1][j] + memo[i][j - 1]
			}
		}
	}
    // fmt.Println(memo)
	return memo[m - 1][n - 1]
}

func uniquePaths4(m int, n int) int {
	memo := make([]int, n)
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				memo[j] = 1
			} else {
				memo[j] += memo[j - 1]
			}
		}
	}
	return memo[n - 1]
}

