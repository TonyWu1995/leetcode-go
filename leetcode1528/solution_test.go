package leetcode1528

import "testing"

func Test_restoreString(t *testing.T) {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}

	res := restoreString(s, indices)

	if res != "leetcode" {
		t.Errorf("no match")
	}
}
