package leetcode20

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	arr := []rune{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			arr = append(arr, v)
		} else {
			lastIndex := len(arr) - 1
			if lastIndex == -1 {
				return false
			}
			lastOpenParenthese := arr[lastIndex]
			if v == m[lastOpenParenthese] {
				arr = arr[:lastIndex]
			} else {
				return false
			}
		}
	}
	return len(arr) == 0
}
