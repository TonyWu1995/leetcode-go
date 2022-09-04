package leetcode2351

func repeatedCharacter(s string) byte {
	arr := make([]int, 26)
	for _, v := range s {
		arr[v-'a']++
		if arr[v-'a'] >= 2 {
			return byte(v)
		}
	}
	return ' '
}
