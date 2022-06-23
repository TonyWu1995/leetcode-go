package leetcode2011

import "testing"

func Test_finalValueAfterOperations(t *testing.T) {

	res := finalValueAfterOperations([]string{"--X", "X++", "X++"})

	if res != 1 {
		t.Errorf("not match 1")
	}
}
