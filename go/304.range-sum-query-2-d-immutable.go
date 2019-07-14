/*
 * @lc app=leetcode id=304 lang=golang
 *
 * [304] Range memo Query 2D - Immutable
 */
 type NumMatrix struct {
    memo [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	mtx := new(NumMatrix)

	m := len(matrix)
    if m == 0 { return *mtx }
	n := len(matrix[0])
	if n == 0 { return *mtx }
	
	mtx.memo = make([][]int, m)
	for i := 0; i < m; i++ {
		mtx.memo[i] = make([]int, n)
	}

	mtx.memo[0][0] = matrix[0][0]
	for c := 1; c < n; c++ {
		mtx.memo[0][c] = mtx.memo[0][c - 1] + matrix[0][c]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				mtx.memo[i][j] = mtx.memo[i - 1][j] + matrix[i][j]
			} else {
				mtx.memo[i][j] = mtx.memo[i - 1][j] + mtx.memo[i][j - 1] - mtx.memo[i - 1][j - 1] + matrix[i][j] 
			}
		}
	}
	
	//fmt.Println(*mtx)
    return *mtx
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.memo[row2][col2] 
	} else if row1 == 0 {
		return this.memo[row2][col2] - this.memo[row2][col1 - 1]
	} else if col1 == 0 {
		return this.memo[row2][col2] - this.memo[row1 - 1][col2]
	} else {
		return this.memo[row2][col2] - this.memo[row1 - 1][col2] - this.memo[row2][col1 - 1] + this.memo[row1 - 1][col1 - 1]
	}
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.memoRegion(row1,col1,row2,col2);
 */

