/*
 * @lc app=leetcode id=464 lang=golang
 *
 * [464] Can I Win
 */

//Time Limit Exceeded: 20 \n 210
func canIWin0(maxChoosableInteger int, desiredTotal int) bool {

	if desiredTotal <= 0 {
		return true
	}

	if (1 + maxChoosableInteger) * maxChoosableInteger / 2 < desiredTotal {
		return false
	}

	nums := make([]int, maxChoosableInteger)
    for i := 1; i <= maxChoosableInteger; i++ {
		nums[i-1] = i
	}

	return playIt(nums, desiredTotal)
}

func playIt(nums []int, desiredTotal int) bool {
	if desiredTotal < 0 {
		return true
	}

	n := len(nums)
	if nums[n - 1] >= desiredTotal {
		return true
	}

	result := false
	for i, v := range(nums) {
		inums := []int{}
		inums = append(append(inums, nums[0: i]...), nums[i+1:]...)
		result = result || !playIt(inums, desiredTotal - v)
	}

	return result
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal <= 0 {
		return true
	}

	if (1 + maxChoosableInteger) * maxChoosableInteger / 2 < desiredTotal {
		return false
	}

	var seq bytes.Buffer
    for i := 1; i <= maxChoosableInteger; i++ {
		seq.WriteString("1")
	}

	memo := make(map[string]bool)

	return playIt(seq.String(), desiredTotal, memo)
}


func playIt(seq string, desiredTotal int, memo map[string]bool) bool {
	if ok, v := memo[seq]; ok {
		return v
	}
	if desiredTotal < 0 {
		return true
	}

	n := len(seq)
	for i := n - 1; i >= 0; i-- {
		if seq[i:i+1] == "1" && i + 1 >= desiredTotal {
			return true
		}
	}


	result := false
	for i, _ := range(seq) {
        iv := seq[i:i+1]
        if iv == "1" {
            iseq := seq[0:i] + "0" + seq[i+1:len(seq)]
            result = result || !playIt(iseq, desiredTotal - (i + 1), memo)
        }
		
	}
	memo[seq] = result

	return result
}
