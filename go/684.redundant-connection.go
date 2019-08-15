/*
 * @lc app=leetcode id=684 lang=golang
 *
 * [684] Redundant Connection
 */
func findRedundantConnection(edges [][]int) []int {
	group := make([]int, len(edges) + 1)
	for _, edge := range(edges) {
		g1 := find(group, edge[0])
		g2 := find(group, edge[1])
		if(g1 == g2) {
			return edge
		}

		group[g1] = g2
	}

	return []int{}
}

func find(group []int, vertex int) int {
	if group[vertex] == 0 { return vertex }
	return find(group, group[vertex])
}

