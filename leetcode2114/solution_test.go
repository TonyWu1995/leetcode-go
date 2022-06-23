package leetcode2114

import (
	"testing"
)

func Test_mostWordsFound(t *testing.T) {
	res := mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"})

	if res != 6 {
		t.Errorf("not match 6")
	}
}
