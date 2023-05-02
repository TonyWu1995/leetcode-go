package leetcode77

func combine(n int, k int) [][]int {
	res := [][]int{}
	arr := make([]int, k)
	i := 0
	for i >= 0 {
		arr[i]++
		if arr[i] > n {
			i--
		} else if i == k-1 {
			res = append(res, append([]int{}, arr...))
		} else {
			i++
			arr[i] = arr[i-1]
		}
	}
	return res
}
