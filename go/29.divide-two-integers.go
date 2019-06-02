/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 *
 * https://leetcode.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (16.14%)
 * Total Accepted:    188.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '10\n3'
 *
 * Given two integers dividend and divisor, divide two integers without using
 * multiplication, division and mod operator.
 * 
 * Return the quotient after dividing dividend by divisor.
 * 
 * The integer division should truncate toward zero.
 * 
 * Example 1:
 * 
 * 
 * Input: dividend = 10, divisor = 3
 * Output: 3
 * 
 * Example 2:
 * 
 * 
 * Input: dividend = 7, divisor = -3
 * Output: -2
 * 
 * Note:
 * 
 * 
 * Both dividend and divisor will be 32-bit signed integers.
 * The divisor will never be 0.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 2^31 − 1 when the
 * division result overflows.
 * 
 * 
 */

 func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	
	if dividend >= 0 && divisor > 0 {
		return dividePositive(dividend, divisor)
	} else if dividend < 0 && divisor < 0 {
		return dividePositive(-dividend, -divisor)
	} else if dividend < 0 {
		return -dividePositive(-dividend, divisor)
	} else {
		return -dividePositive(dividend, -divisor)
	} 
}

func dividePositive(dividend int, divisor int) int {
	times := 0
	sum := 0
	currSum := divisor
	currTimes := 1
	currDividend := dividend
	for sum <= dividend {

		if currSum > currDividend {
			break
		} else if currSum + currSum > currDividend {
			sum += currSum
			times += currTimes
			currDividend -= currSum

			currSum = divisor
			currTimes = 1
		} else {
			currSum = currSum + currSum
			currTimes = currTimes + currTimes
		}
	}

	if times > int(math.Pow(2, 31) - 1) {
		return int(math.Pow(2, 31) - 1)
	} else {
		return times
	}
	
}

