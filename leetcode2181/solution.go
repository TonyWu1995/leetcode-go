package leetcode2181

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	if head == nil || head.Val == 0 && head.Next == nil {
		return nil
	}
	curr := head
	mod := &ListNode{}
	for curr.Next != nil {
		if curr.Val == 0 {
			mod = curr
			curr = curr.Next
		}
		if curr.Val != 0 {
			mod.Val += curr.Val
			*curr = *curr.Next
		}
	}
	mod.Next = nil
	return head
}
