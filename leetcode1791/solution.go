package leetcode1791

func findCenter(edges [][]int) int {
	m := map[int]int{}
	for _, v := range edges {
		a := v[0]
		b := v[1]
		m[a]++
		m[b]++
	}
	res := 0
	for k, v := range m {
		if max(m[res], m[k]) == v {
			res = k
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
