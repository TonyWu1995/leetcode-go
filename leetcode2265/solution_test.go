package leetcode2265

import (
	"testing"
)

func Test_averageOfSubtree_case1(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	res := averageOfSubtree(root)

	if res != 1 {
		t.Errorf("not match")
	}
}

func Test_averageOfSubtree_case2(t *testing.T) {

	leafNode1 := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	leafNode2 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	leafNode6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}

	interNode1 := &TreeNode{
		Val:   8,
		Left:  leafNode1,
		Right: leafNode2,
	}
	interNode2 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: leafNode6,
	}
	root := &TreeNode{
		Val:   4,
		Left:  interNode1,
		Right: interNode2,
	}
	//[4,8,5,0,1,null,6]
	res := averageOfSubtree(root)
	if res != 5 {
		t.Errorf("not match")
	}
}
