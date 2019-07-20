/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
func numDecodings0(s string) int {
	n := len(s)
	if n == 0 { return 0 }
	return decode(s)
}

func decode(s string) int {
	n := len(s)
	if n == 0 { return 1 }
	if n == 1 { 
		return validate(s)
	}
	if s[0:1] == "0" {
		return 0
	}
	var t1, t2 int
	if validate(s[0:1]) == 1 {
		t1 := decode(s[1:]) 
	}
	if validate(s[0:2]) == 1 {
		t2 := decode(s[2:])
	}
	
	return t1 + t2
}

func validate(str string) int {
	num, _ := strconv.Atoi(str)
	if str[0:1] == "0" {
		return 0
	} else if num > 26 || num < 1 {
		return 0
	} else {
		return 1
	}
}

func numDecodings(s string) int {
	n := len(s)
	if n == 0 { return 0 }
	return decode(s)
}

