package leetcode2011

func finalValueAfterOperations(operations []string) int {
	res := 0
	for _, v := range operations {
		if v[1] == '-' {
			res--
		} else {
			res++
		}
	}
	return res
}
