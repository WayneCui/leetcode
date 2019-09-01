/*
 * @lc app=leetcode id=1130 lang=golang
 *
 * [1130] Minimum Cost Tree From Leaf Values
 */
//Time Limit Exceeded
func mctFromLeafValues(arr []int) int {
	n := len(arr)
	if n == 0 { return 0 }
	if n == 1 { return 0 }
	if n == 2 { return arr[0] * arr[1] }

	minsum, _ := helper(arr)
	return minsum
}

func helper(arr []int) (minSum int, maxLeaf int) {
	n := len(arr)
	if n == 1 { return 0, arr[0]}
	if n == 2 { return arr[0] * arr[1], max(arr[0], arr[1])}

	ms := math.MaxInt32
	ml := 0
	for i := 1; i < n; i++ {
		sum1, ml1 := helper(arr[0:i])
		sum2, ml2 := helper(arr[i:])
		sum := ml1 * ml2 + sum1 + sum2

		ms = min(ms, sum)
		ml = max(ml, ml1, ml2)
	}

	return ms, ml
}

func serialize(s []int) string {
	if len(s) == 0 { return "" }

	res := ""
	for _, v := range s {
		res += strconv.Itoa(v) + ","
	}

	return res
}

//with memo, 264 ms
func mctFromLeafValues(arr []int) int {
	n := len(arr)
	if n == 0 { return 0 }
	if n == 1 { return 0 }
	if n == 2 { return arr[0] * arr[1] }
	
	memo := make(map[string][]int)
	
	minsum, _ := helper(arr, memo)
	return minsum
}

func helper(arr []int, memo map[string][]int) (minSum int, maxLeaf int) {
	n := len(arr)
	if n == 1 { return 0, arr[0]}
	if n == 2 { return arr[0] * arr[1], max(arr[0], arr[1])}

	ms := math.MaxInt32
	ml := 0
	for i := 1; i < n; i++ {
		sum1, ml1 := 0, 0
		if v, ok := memo[serialize(arr[0:i])]; ok {
			sum1, ml1 = v[0], v[1]
		} else {
			sum1, ml1 = helper(arr[0:i], memo)
			memo[serialize(arr[0:i])] = []int{sum1, ml1}
		}

		sum2, ml2 := 0, 0
		if v, ok := memo[serialize(arr[i:])]; ok {
			sum2, ml2 = v[0], v[1]
		} else {
			sum2, ml2 = helper(arr[i:], memo)
            memo[serialize(arr[i:])] = []int{sum2, ml2}
		}

		sum := ml1 * ml2 + sum1 + sum2

		ms = min(ms, sum)
		ml = max(ml, ml1, ml2)
	}

	return ms, ml
}

//dynamic programming
func mctFromLeafValues(arr []int) int {
	n := len(arr)
	if n == 0 { return 0 }
	if n == 1 { return 0 }
	if n == 2 { return arr[0] * arr[1] }
	
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = []int{0, 0}
        }
	}

	for i := 0; i < n; i++ {
		dp[i][i] = []int{0, arr[i]}
		if i + 1 < n {
			dp[i][i + 1] = []int{arr[i] * arr[i + 1], max(arr[i], arr[i + 1])}
		}
	}

	for i := 2; i < n; i++ { // step
		for j := 0; j < n; j++ { // row
			if j + i < n {
				sum := math.MaxInt32
				for k := j; k < i + j; k++ { //split
					sum = min(sum, dp[j][k][1] * dp[k+1][i+j][1] + dp[j][k][0] + dp[k+1][i+j][0])
				}
                mx := max(arr[j:j + i + 1]...)
				dp[j][j + i] = []int{sum, mx}
			}
		}
	}

	return dp[0][n - 1][0]
}

func max(a...int) int {
	m := math.MinInt32
	for _, v := range a {
		if v > m {
			m = v
		}
	}

	return m
}

func min(a...int) int {
	m := math.MaxInt32
	for _, v := range a {
		if v < m {
			m = v
		}
	}

	return m
}


