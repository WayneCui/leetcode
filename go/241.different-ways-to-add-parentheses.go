/*
 * @lc app=leetcode id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 */
func diffWaysToCompute(input string) []int {
	memo := make(map[string][]int)
	return compute(input, memo)
}
func compute(input string, memo map[string][]int) []int {
	if v, ok := memo[input]; ok {
		return v
	}

	if !strings.Contains(input, "+") && 
		!strings.Contains(input, "-") && 
		!strings.Contains(input, "*") {
			i, _ := strconv.Atoi(input)
			return []int{ i }
	}

	output := []int{}
    for i, s := range strings.Split(input, "") {
		left := compute(input[0:i], memo)
		right := compute(input[i+1:], memo)
		switch s {
			case "+":
				for _, v := range left {
					for _, w := range right {
						output = append(output, v + w)
					}
				}
			case "-":
				for _, v := range left {
					for _, w := range right {
						output = append(output, v - w)
					}
				}
			case "*":
				for _, v := range left {
					for _, w := range right {
						output = append(output, v * w)
					}
				}
			default:
		}
	}

	memo[input] = output
	return output
}

