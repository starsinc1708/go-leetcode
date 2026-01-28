package p515_findlargestvalineachtreerow

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		level_size := len(q)
		curLevelMax := math.MinInt64
		for range level_size {
			node := q[0]
			q = q[1:]
			curLevelMax = max(curLevelMax, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		result = append(result, curLevelMax)
	}

	return result
}
