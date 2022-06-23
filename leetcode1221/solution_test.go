package leetcode1221

import "testing"

func Test_balancedStringSplit(t *testing.T) {
	s := "RLRRLLRLRL"

	res := balancedStringSplit(s)

	if res != 4 {
		t.Errorf("no match")
	}

}
