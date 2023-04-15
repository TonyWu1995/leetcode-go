package leetcode1646

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	res := 0
	for i := 2; i < n+1; i++ {
		number := nums[i/2]
		if i%2 == 1 {
			number += nums[i/2+1]
		}
		nums[i] = number
		if res < number {
			res = number
		}
	}
	return res
}
