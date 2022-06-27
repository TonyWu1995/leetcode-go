package leetcode1282

func groupThePeople(groupSizes []int) [][]int {
	res := make([][]int, 0)
	temp := make(map[int][]int)

	for k, v := range groupSizes {
		_, ok := temp[v]
		if !ok {
			temp[v] = []int{k}
		} else {
			temp[v] = append(temp[v], k)
		}
		if v == len(temp[v]) {
			res = append(res, temp[v])
			temp[v] = nil
		}
	}
	return res
}
