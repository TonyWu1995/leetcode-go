package leetcode1528

func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for k, v := range indices {
		res[v] = s[k]
	}
	return string(res)
}
