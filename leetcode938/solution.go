package leetcode938

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	if root == nil {
		return res
	}
	if root.Val >= low && root.Val <= high {
		res += root.Val
	}
	res += rangeSumBST(root.Left, low, high)
	res += rangeSumBST(root.Right, low, high)
	return res
}
