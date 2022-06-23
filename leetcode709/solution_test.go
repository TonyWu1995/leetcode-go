package leetcode709

import "testing"

func Test_toLowerCase(t *testing.T) {
	s := "Hello"

	res := toLowerCase(s)

	if res != "hello" {
		t.Errorf("not match")
	}

}
