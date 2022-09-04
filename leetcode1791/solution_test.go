package leetcode1791

import "testing"

func Test_findCenter(t *testing.T) {

	input := [][]int{
		{1, 2},
		{2, 3},
		{4, 2},
	}

	res := findCenter(input)

	if res != 2 {
		t.Errorf("not match")
	}

}
