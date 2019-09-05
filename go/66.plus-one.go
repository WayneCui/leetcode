/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
 func plusOne(digits []int) []int {
    n := len(digits)
    carry := 1
    for i := n - 1; i >= 0; i-- {
        r := digits[i] + carry
        if r > 9 {
            digits[i] = r - 10
            carry = 1
        } else {
            digits[i] = r
            carry = 0
        }
    }
    
    if carry == 1 {
        digits = insert(digits, 0, 1)
    }
    
    return digits
}

func insert(a []int, i int, v int) []int {
    a = append(a, 0)
    copy(a[1:], a[0:])
    a[0] = v
    return a
}

