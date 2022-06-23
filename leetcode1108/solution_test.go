package leetcode1108

import "testing"

func Test_defangIPaddr(t *testing.T) {
	address := "1.1.1.1"
	res := defangIPaddr(address)

	if res != "1[.]1[.]1[.]1" {
		t.Errorf("not match")
	}

}
