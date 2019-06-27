/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
 func longestPalindrome(s string) string {
	n := len(s)
	str := []rune(s)
    if n == 0 || n == 1 {
        return s
    }
	
	var from int
	var to int
	longest := 0
    for i := 0; i < n; i++ {		
		if i + 2 < n && str[i] == str[i + 2] {
            length := 1
            end := i + 2
            begin := i
			for begin >= 0 && end < n {
				if str[begin] != str[end] {
					break
				}

				length += 2
				begin -= 1
				end += 1
			}
            
            if longest < length {
                longest = length
                from = begin
                to = end
            }
		} 
        
        if i + 1 < n && str[i] == str[i + 1] {
            length := 0
            begin := i
            end := i + 1
			for begin >= 0 && end < n {
				if str[begin] != str[end] {
					break
				}

				length += 2
				begin -= 1
				end += 1
			}
            
            if longest < length {
                longest = length
                from = begin
                to = end
            }
		} 

		
	}

    // fmt.Println(strconv.Itoa(from + 1) + ":" + strconv.Itoa(to))
    if (from + 1) > to {
        return string(str[0])
    } else {
        return string(str[(from + 1) : to ])
    }
    
}

