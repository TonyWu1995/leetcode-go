package leetcode771

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	jewels := "aA"
	stones := "aAAbbbb"

	res := numJewelsInStones(jewels, stones)

	if res != 3 {
		t.Errorf("not match")
	}
}
