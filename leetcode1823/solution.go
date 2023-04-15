package leetcode1823

func findTheWinner(n int, k int) int {
	arr := []int{}
	for i := 1; i < n+1; i++ {
		arr = append(arr, i)
	}
	s := 0
	for len(arr) != 1 {
		rem := (s + k - 1) % len(arr)
		arr = append(arr[:rem], arr[rem+1:]...)
		s = rem
	}
	return arr[0]
}
