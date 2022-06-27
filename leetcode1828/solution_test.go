package leetcode1828

import (
	"testing"
)

func Test_countPoints(t *testing.T) {
	points := [][]int{
		{1, 3},
		{3, 3},
		{5, 3},
		{2, 2},
	}

	query := [][]int{
		{2, 3, 1},
		{4, 3, 1},
		{1, 1, 2},
	}

	res := countPoints(points, query)

	if !Equal(res, []int{3, 2, 2}) {
		t.Errorf("no match")
	}

}
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
