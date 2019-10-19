/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func reverseWords0(s string) string {
	low := 0
	high := 0
	idx := 0
    n := len(s)
	bytes := []byte(s)
    for i := 0; i < n;  {
        fmt.Println(idx)
		for idx < n && bytes[idx] != ' ' {
			idx += 1
		}

        low = i
		high = idx - 1
		for ; low < high; low, high = low + 1, high - 1 {
			bytes[low], bytes[high] = bytes[high], bytes[low]
		}
        
        idx += 1
        i = idx
	}

	return string(bytes)
}

func reverseWords(s string) string {
	bytes := []byte(s)
    bytes = append(bytes, ' ')
	prev := -1
    n := len(bytes)

	for i := 0; i < n; i++ {
		if bytes[i] == ' ' {
			low := prev + 1
			high := i - 1
			for ; low < high; low, high = low+1, high-1 {
				bytes[low], bytes[high] = bytes[high], bytes[low]
			}
            prev = i
		}
	}

    return string(bytes[0:n-1])
}
// @lc code=end

