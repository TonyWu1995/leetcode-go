package leetcode977

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	n := len(nums)
	left, right := 0, n-1
	for i := n - 1; i >= 0; i-- {
		square := 0
		if Abs(nums[left]) < Abs(nums[right]) {
			square = nums[right]
			right--
		} else {
			square = nums[left]
			left++
		}
		result[i] = square * square
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
