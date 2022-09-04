package leetcode2341

func numberOfPairs(nums []int) []int {
	m := map[int]int{}
	size := len(nums)
	count := 0
	for _, v := range nums {
		m[v]++
		if m[v] == 2 {
			size -= 2
			count++
			m[v] = 0
		}
	}
	return []int{count, size}
}
