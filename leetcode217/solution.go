package leetcode217

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, v := range nums {
		if _, exist := m[v]; exist {
			return true
		}
		m[v] = true
	}
	return false
}
