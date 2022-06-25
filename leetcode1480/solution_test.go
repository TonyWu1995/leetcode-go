package leetcode1480

import (
	"testing"
)

func Test_runningSum(t *testing.T) {
	input := []int{1, 2, 3, 4}

	res := runningSum(input)

	expect := []int{1, 3, 6, 10}

	if !Equal(res, expect) {
		t.Errorf("not match")
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
