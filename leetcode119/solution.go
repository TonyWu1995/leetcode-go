package Leetcode120

func getRow(rowIndex int) []int {
	pascal := Triangle(rowIndex + 1)
	return pascal[rowIndex]
}

func Triangle(depth int) [][]int {
	if depth == 1 {
		return [][]int{{1}}
	}
	triangle := make([][]int, depth)
	for i, row := range triangle {
		for j := 0; j < i+1; j++ {
			row = append(row, 1)
		}
		triangle[i] = row
	}
	for i, row := range triangle[2:] {
		for j := 1; j < len(row)-1; j++ {
			row[j] = triangle[i+1][j-1] + triangle[i+1][j]
		}
	}
	return triangle
}
