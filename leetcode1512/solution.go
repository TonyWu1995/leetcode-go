package leetcode1512

func numIdenticalPairs(nums []int) int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
	}
	result := 0
	for _, val := range dict {
		result += val * (val - 1) / 2
	}
	return result
}
