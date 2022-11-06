package leetcode2221

func triangularSum(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	idx := n - 1
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			nums[j] = (nums[j] + nums[j+1]) % 10
		}
		nums[idx] = 0
		idx--
	}
	return nums[0]
}
