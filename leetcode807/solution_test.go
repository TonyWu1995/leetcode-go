package leetcode807

import "testing"

func Test_maxIncreaseKeepingSkyline(t *testing.T) {
	input := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}

	res := maxIncreaseKeepingSkyline(input)

	if res != 35 {
		t.Errorf("no match")
	}

}
