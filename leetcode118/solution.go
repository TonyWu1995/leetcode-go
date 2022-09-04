package Leetcode120

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	triangle := make([][]int, numRows)
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
