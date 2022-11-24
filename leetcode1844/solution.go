package leetcode1844

func replaceDigits(s string) string {
	arr := []byte(s)
	for i := 1; i < len(s); i += 2 {
		arr[i] = (arr[i-1] - 'a') + arr[i] - '0' + 'a'
	}
	return string(arr)
}
