package p98_validate_bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeRange struct {
	Node *TreeNode
	Min  int
	Max  int
}

func isValidBST(root *TreeNode) bool {
	if root.Right == nil && root.Left == nil {
		return true
	}

	stack := make([]*NodeRange, 0)
	stack = append(stack, &NodeRange{
		root,
		math.MinInt64,
		math.MaxInt64,
	})

	for len(stack) > 0 {
		node := stack[len(stack)-1].Node
		minR := stack[len(stack)-1].Min
		maxR := stack[len(stack)-1].Max

		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}

		if node.Val <= minR || node.Val >= maxR {
			return false
		}

		stack = append(stack, &NodeRange{
			node.Left,
			minR,
			node.Val,
		})

		stack = append(stack, &NodeRange{
			node.Right,
			node.Val,
			maxR,
		})
	}

	return true
}
