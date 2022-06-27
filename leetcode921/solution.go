package leetcode921

func minAddToMakeValid(s string) int {
	stack := 0
	res := 0
	for _, v := range s {
		if v == 40 {
			stack += 1
		}
		if v == 41 {
			if stack == 0 {
				res += 1
			} else {
				stack -= 1
			}
		}
	}
	return res + stack
}
