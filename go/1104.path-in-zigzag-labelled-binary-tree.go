/*
 * @lc app=leetcode id=1104 lang=golang
 *
 * [1104] Path In Zigzag Labelled Binary Tree
 */
func pathInZigZagTree(label int) []int {
	if label == 1 {
		return []int{1}
	}
	// 1. 
	tmp := label
	layer := 1
	count := 0
	for tmp > 0 {
		count = pow2(layer - 1)
		if(tmp <= count) { break }
		tmp = tmp - count
		layer += 1
	}

	// 2.
	pos := 0
	if isEven(layer) {
		pos = pow2(layer - 1) - tmp
	} else {
		pos = tmp - 1
	}

	// 3.
	path := []int{label}

	for i := layer - 1; i > 0; i-- {
		min := pow2(i - 1)
		max := pow2(i) - 1

		pos = pos / 2
		if isEven(i) {
			path = append(path, max - pos)
		} else {
			path = append(path, min + pos)
		}

	}

	sort.Ints(path)
	return path
}

func isEven(n int) bool {
	return n % 2 == 0
}

func pow2(n int) int {
	return int(math.Pow(2, float64(n)))
}
