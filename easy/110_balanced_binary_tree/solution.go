package p110_balanced_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(max(height(root.Left), height(root.Right)))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lHeight := height(root.Left)
	rHeight := height(root.Right)

	if math.Abs(float64(lHeight-rHeight)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}
