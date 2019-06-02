/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (34.74%)
 * Total Accepted:    215.4K
 * Total Submissions: 619.6K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 * 
 * 
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * Output: false
 * 
 */
 func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	low := 0
	high := m - 1
	mid := 0
	targetRow := 0
	for low < high {
		mid = (low + high) / 2
		
		if target > matrix[mid][n - 1] {
			low = mid + 1
		} else if target < matrix[mid][0] {
			high = mid
		} else {
			low = mid
			break
		}
	}

	targetRow = low
	fmt.Println(targetRow)
	low = 0
	high = n - 1
	mid = 0
	for low < high {
		mid = (low + high) / 2
		if target > matrix[targetRow][mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if target == matrix[targetRow][low] {
		return true
	} else {
		return false
	}
}

