package leetcode2181

import (
	"fmt"
	"testing"
)

func Test_mergeNodes(t *testing.T) {

	node6 := &ListNode{
		0,
		nil,
	}
	node5 := &ListNode{
		5,
		node6,
	}
	node4 := &ListNode{
		4,
		node5,
	}
	node3 := &ListNode{
		0,
		node4,
	}

	node2 := &ListNode{
		1,
		node3,
	}

	node1 := &ListNode{
		3,
		node2,
	}

	head := &ListNode{
		0,
		node1,
	}

	res := mergeNodes(head)
	for res.Next != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
