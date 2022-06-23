package leetcode1221

func balancedStringSplit(s string) int {
	res := 0
	temp := 0
	for _, v := range s {
		if v == 82 {
			temp++
		} else {
			temp--
		}
		if temp == 0 {
			res++
		}
	}
	return res
}
