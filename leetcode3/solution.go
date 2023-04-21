package leetcode3

//abcabcbb
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	start := 0
	m := make(map[string]int)
	l := 0
	for end := 0; end < len(s); end++ {
		if val, ok := m[string(s[end])]; ok {
			start = max(start, val+1)
		}
		m[string(s[end])] = end
		l = max(l, end-start+1)
	}
	return l
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
