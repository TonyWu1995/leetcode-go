package leetcode2363

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := map[int]int{}
	for _, v := range items1 {
		m[v[0]] = v[1]
	}
	for _, v := range items2 {
		m[v[0]] += v[1]
	}
	res := [][]int{}
	for k, v := range m {
		res = append(res, []int{k, v})
	}
	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })
	return res
}
