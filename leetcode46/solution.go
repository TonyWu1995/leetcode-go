package leetcode46

func permute(nums []int) [][]int {
	var res [][]int
	helper(&res, nums, []int{})
	return res
}

func helper(res *[][]int, nums []int, path []int) {
	if len(nums) == 0 {
		*res = append(*res, path)
	}
	for index, num := range nums {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		helper(res, append(numsCopy[:index], numsCopy[index+1:]...), append(pathCopy, num))
	}
}
