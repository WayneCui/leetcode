func countBinarySubstrings0(s string) int {
	splits := []int{0}
	n := len(s)
	prev := s[0]
	for i := 1; i < n; i++ {
		if s[i] != prev {
			splits = append(splits, i)
		}
		prev = s[i]
	}
	splits = append(splits, n)

	result := 0
	for i := 0; i < len(splits) - 2; i++ {
		result += min(splits[i + 1] - splits[i], splits[i + 2] - splits[i + 1])
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func countBinarySubstrings1(s string) int {
	s = strings.Replace(s, "01", "0 1", -1)
	s = strings.Replace(s, "10", "1 0", -1)
	parts := strings.Split(s, " ")

	result := 0
	for i := 0; i < len(parts) - 1; i++ {
		result += min(len(parts[i]), len(parts[i+1]))
	}

	return result
}

func countBinarySubstrings(s string) int {
	pre, cur, result := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i - 1] {
			cur += 1
		} else {
			result += min(cur, pre)
			pre = cur
			cur = 1
		}
	}

	return result += min(cur, pre)
}