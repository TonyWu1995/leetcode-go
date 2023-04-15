package leetcode2085

func countWords(words1 []string, words2 []string) int {
	m := map[string]int{}
	for _, v := range words1 {
		if _, isExist := m[v]; !isExist {
			m[v] = 1
		} else {
			m[v] = 0
		}
	}
	count := 0
	for _, v := range words2 {
		if val, isExist := m[v]; isExist {
			if val == 1 {
				count++
				m[v]++
			}
			if val == 2 {
				count--
				m[v]++
			}
		}
	}
	return count
}
