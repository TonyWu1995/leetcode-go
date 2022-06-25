package leetcode1859

import "testing"

func Test_sortSentence(t *testing.T) {
	s := "is2 sentence4 This1 a3"

	res := sortSentence(s)

	if res != "This is a sentence" {
		t.Errorf("no match")
	}

}
