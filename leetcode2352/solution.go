package leetcode2352

import (
	"fmt"
	"strconv"
	"strings"
)

func join(arr []int) string {
	tmp := []string{}
	for i := range arr {
		text := strconv.Itoa(arr[i])
		tmp = append(tmp, text)
	}
	return strings.Join(tmp, "")
}

func equalPairs(grid [][]int) int {
	cols := make(map[string]int)
	rows := make(map[string]int)
	n := len(grid) // n x n matrix
	for i := 0; i < n; i++ {
		var row, col strings.Builder
		for j := 0; j < n; j++ {
			row.WriteString(fmt.Sprintf("%d;", grid[i][j]))
			col.WriteString(fmt.Sprintf("%d;", grid[j][i]))
		}
		cols[col.String()]++
		rows[row.String()]++
	}
	var res int
	for col, count1 := range cols {
		if count2, ok := rows[col]; ok {
			res += count1 * count2
		}
	}
	return res
}
