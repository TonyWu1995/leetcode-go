package leetcode938

import "testing"

func Test_rangeSumBST(t *testing.T) {
	leafNode18 := &TreeNode{Val: 18, Left: nil, Right: nil}
	interNode15 := &TreeNode{Val: 15, Left: nil, Right: leafNode18}
	leafNode7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	leafNode3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	interNode5 := &TreeNode{Val: 5, Left: leafNode3, Right: leafNode7}
	rootNode := &TreeNode{Val: 10, Left: interNode5, Right: interNode15}

	res := rangeSumBST(rootNode, 7, 15)

	if res != 32 {
		t.Errorf("not match 32")
	}
}
