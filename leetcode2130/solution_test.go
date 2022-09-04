package leetcode2130

import "testing"

func Test_pairSum(t *testing.T) {
	fourth := &ListNode{
		1,
		nil,
	}
	third := &ListNode{
		2,
		fourth,
	}
	second := &ListNode{
		4,
		third,
	}
	head := &ListNode{
		5,
		second,
	}

	res := pairSum(head)

	if res != 6 {
		t.Errorf("not match")
	}

}
