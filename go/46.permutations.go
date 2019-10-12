/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute0(nums []int) [][]int {
    output := [][]int{}

    if len(nums) == 1 {
        output = append(output, []int{nums[0]})
        return output
    }
    for i, v := range nums {
        for _, part := range permute(del(nums, i)) {
            output = append(output, append([]int{v}, part...))
        }
    }
                                     
    return output
}

func del(nums []int, i int) []int {
    tmp := append(nums[:0:0], nums...)
    return append(tmp[:i], tmp[i + 1:]...)  
}

// backtrack with marks
func permute(nums []int) [][]int {
    output := make([][]int, 0)
    tmpList := make([]int, 0)
    marks := make([]bool, len(nums))
	backtrack(&output, tmpList, nums, marks)

	return output
}

func backtrack(output *[][]int, tmpList []int, nums []int, marks []bool) {
	if len(tmpList) == len(nums) {
        tmp := make([]int, len(tmpList))
        copy(tmp, tmpList)
		*output = append(*output, tmp)
		return
	} 

	for i, v := range nums {
		if marks[i] {
			continue
		}

		tmpList = append(tmpList, v)
        marks[i] = true
		backtrack(output, tmpList, nums, marks)
        tmpList = tmpList[:len(tmpList) - 1]
        marks[i] =false
	}
}

// backtrack
func permute(nums []int) [][]int {
    output := make([][]int, 0)
    tmpList := make([]int, 0)
	backtrack(&output, tmpList, nums)

	return output
}

func backtrack(output *[][]int, tmpList []int, nums []int) {
	if len(tmpList) == len(nums) {
        tmp := make([]int, len(tmpList))
        copy(tmp, tmpList)
		*output = append(*output, tmp)
		return
	} 

	for _, v := range nums {
		if contains(tmpList, v) {
			continue
		}

		tmpList = append(tmpList, v)
		backtrack(output, tmpList, nums)
        tmpList = tmpList[:len(tmpList) - 1]
	}
}

func contains(a []int, num int) bool {
	for _, v := range a {
		if v == num {
			return true
		}
	}

	return false
}





// @lc code=end

