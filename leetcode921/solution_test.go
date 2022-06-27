package leetcode921

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	s := "())"

	res := minAddToMakeValid(s)

	if res != 1 {
		t.Errorf("not match")
	}
}

func Test_minAddToMakeValid_case2(t *testing.T) {
	s := "((("

	res := minAddToMakeValid(s)

	if res != 3 {
		t.Errorf("not match")
	}
}

func Test_minAddToMakeValid_case3(t *testing.T) {
	s := "()))(("

	res := minAddToMakeValid(s)

	if res != 4 {
		t.Errorf("not match")
	}
}
