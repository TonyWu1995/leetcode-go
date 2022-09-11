package leetcode2194

func cellsInRange(s string) []string {
	res := []string{}
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			res = append(res, string([]byte{i, j}))
		}
	}
	return res
}
