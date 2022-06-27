package leetcode1347

func minSteps(s string, t string) int {
	temp := make([]int, 26)
	for i := 0; i < len(s); i++ {
		temp[s[i]-'a'] += 1
		temp[t[i]-'a'] -= 1
	}
	res := 0
	for i := range temp {
		if temp[i] > 0 {
			res += temp[i]
		}
	}

	return res
}
