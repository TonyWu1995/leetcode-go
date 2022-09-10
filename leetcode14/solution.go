package leetcode14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for index, val := range strs[0] {
		for _, j := range strs[1:] {
			if index == len(j) || byte(val) != j[index] {
				return strs[0][:index]
			}
		}
	}
	return strs[0]
}
