package leetcode1672

import "testing"

func Test_maximumWealth_case1(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}

	res := maximumWealth(input)

	if res != 6 {
		t.Errorf("not match")
	}
}

func Test_maximumWealth_case2(t *testing.T) {
	input := [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}

	res := maximumWealth(input)

	if res != 10 {
		t.Errorf("not match")
	}
}
