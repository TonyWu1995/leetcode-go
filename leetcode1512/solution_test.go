package leetcode1512

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
	nums := []int{1, 2, 3, 1, 1, 3}
	result := numIdenticalPairs(nums)

	if result != 4 {
		t.Errorf("not match 4")
	}
}
