package leetcode1828

func countPoints(points [][]int, queries [][]int) []int {
	size := len(queries)
	res := make([]int, len(queries))
	for i := 0; i < size; i++ {
		x := queries[i][0]
		y := queries[i][1]
		r := queries[i][2]
		for j := 0; j < len(points); j++ {
			px := points[j][0]
			py := points[j][1]
			if (px-x)*(px-x)+(py-y)*(py-y) <= r*r {
				res[i]++
			}
		}
	}
	return res
}
