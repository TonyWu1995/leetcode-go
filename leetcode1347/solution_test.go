package leetcode1347

import "testing"

func Test_minSteps(t *testing.T) {
	s := "bab"
	input2 := "aba"
	res := minSteps(s, input2)

	if res != 1 {
		t.Errorf("not match")
	}
}

func Test_minSteps_case2(t *testing.T) {
	s := "leetcode"
	input2 := "practice"
	res := minSteps(s, input2)

	if res != 5 {
		t.Errorf("not match")
	}
}

func Test_minSteps_case3(t *testing.T) {
	s := "anagram"
	input2 := "mangaar"
	res := minSteps(s, input2)

	if res != 0 {
		t.Errorf("not match")
	}
}
