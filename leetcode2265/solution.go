package leetcode2265

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	_, _, res := sumAndCountOfSubTree(root)
	return res
}

func sumAndCountOfSubTree(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}
	lSum, lCount, lAverage := sumAndCountOfSubTree(root.Left)
	rSum, rCount, rAverage := sumAndCountOfSubTree(root.Right)

	sum := rSum + lSum + root.Val
	count := rCount + lCount + 1
	average := lAverage + rAverage
	if root.Val == sum/count {
		average += 1
	}
	return sum, count, average
}
