package leetcode2130

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	tmp := []int{}
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	res := 0
	size := len(tmp) - 1
	for i, j := 0, size; i < j; i, j = i+1, j-1 {
		res = max(res, tmp[i]+tmp[j])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
