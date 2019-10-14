/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	var s1 []byte
	var s2 []byte
    padding := []byte{}
    for i := 0; i < abs(n1 - n2); i++ {
        padding = append(padding, '0')
    }

	if n1 > n2 {
		s1 = []byte(num1)
		s2 = append(padding, []byte(num2)...)
	} else {
		s1 = []byte(num2)
		s2 = append(padding, []byte(num1)...)
	}
    fmt.Println(s1)
    fmt.Println(s2)
    
    var carry byte
    for i := len(s1) - 1; i >= 0; i-- {
		digit1 := (s1[i] - '0')
		digit2 := (s2[i] - '0')
		if digit1 + digit2 + carry < 10 {
			s1[i] = digit1 + digit2 + carry + '0'
			carry = 0
		} else {
			s1[i] = digit1 + digit2 + carry - 10 + '0'
			carry = 1
		}
	}
    
    if(carry == 1) {
        s1 = append([]byte{'1'}, s1...)
    }

	return string(s1)
}

func abs(a int) int {
    if(a >= 0) {
        return a
    } else {
        return -a
    }
}

// @lc code=end

